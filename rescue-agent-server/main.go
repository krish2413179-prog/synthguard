package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum" // Required for CallMsg
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/machinebox/graphql"
)

/* ================= GLOBAL STATE ================= */

var (
	SignatureStore = make(map[string]SignatureData)
	storeMutex     sync.RWMutex

	ethClient *ethclient.Client
	gqlClient *graphql.Client

	agentContractAddr common.Address
	agentPrivateKey   string

	// Address of the MockLending contract (From your logs)
	lendingContractAddr = common.HexToAddress("0xa0f95A73BA2c1395E9F4B95e6F6b7faF3E07A447")
)

/* ================= DATA STRUCTS ================= */

type SignatureData struct {
	UserAddress string   `json:"owner"`
	Value       *big.Int `json:"value"`
	Deadline    *big.Int `json:"deadline"`
	V           uint8    `json:"v"`
	R           [32]byte `json:"r"`
	S           [32]byte `json:"s"`
}

type CrisisUser struct {
	User string `json:"user"`
}

/* ================= CORS ================= */

func enableCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

/* ================= HTTP HANDLER ================= */

func signatureHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}

	var payload struct {
		Owner    string `json:"owner"`
		Value    string `json:"value"`
		Deadline string `json:"deadline"`
		V        uint8  `json:"v"`
		R        string `json:"r"`
		S        string `json:"s"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	value, _ := new(big.Int).SetString(payload.Value, 10)
	deadline, _ := new(big.Int).SetString(payload.Deadline, 10)

	rBytes, _ := hex.DecodeString(strings.TrimPrefix(payload.R, "0x"))
	sBytes, _ := hex.DecodeString(strings.TrimPrefix(payload.S, "0x"))

	var rArr, sArr [32]byte
	copy(rArr[:], rBytes)
	copy(sArr[:], sBytes)

	storeMutex.Lock()
	SignatureStore[payload.Owner] = SignatureData{
		UserAddress: payload.Owner,
		Value:       value,
		Deadline:    deadline,
		V:           payload.V,
		R:           rArr,
		S:           sArr,
	}
	storeMutex.Unlock()

	log.Println("Stored signature for", payload.Owner)
	w.WriteHeader(http.StatusOK)
}

/* ================= RESCUE EXECUTION ================= */

func executeRescue(parsedABI abi.ABI, borrower string, sig SignatureData) error {
	// --- STEP 1: Get the Actual Debt on-chain ---

	// Minimal ABI for the debt mapping: function debt(address) returns (uint256)
	const lendingABIJSON = `[{"inputs":[{"internalType":"address","name":"user","type":"address"}],"name":"debt","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`

	lendingParsedABI, _ := abi.JSON(strings.NewReader(lendingABIJSON))

	// Pack the call to 'debt(borrower)'
	dataDebt, err := lendingParsedABI.Pack("debt", common.HexToAddress(borrower))
	if err != nil {
		return err
	}

	// Call the contract (Static Call, no gas spent)
	msg := ethereum.CallMsg{
		To:   &lendingContractAddr,
		Data: dataDebt,
	}
	result, err := ethClient.CallContract(context.Background(), msg, nil)
	if err != nil {
		return err
	}

	// Unpack the result
	var out []interface{}
	out, err = lendingParsedABI.Unpack("debt", result)
	if err != nil {
		return err
	}
	actualDebt := out[0].(*big.Int)
	log.Println("🔍 On-Chain Debt Check:", actualDebt.String(), "| Signed Permit:", sig.Value.String())

	// --- STEP 2: Calculate Repay Amount ---

	repayAmount := new(big.Int)

	// If Debt < Permit Amount, only pay the Debt
	if actualDebt.Cmp(sig.Value) < 0 {
		repayAmount = actualDebt
		log.Println("⚠️  Debt is smaller than Permit. Repaying exact debt:", repayAmount)
	} else {
		// If Debt >= Permit Amount, pay the full Permit Amount
		repayAmount = sig.Value
		log.Println("✅  Permit covers full or partial debt. Repaying signed amount:", repayAmount)
	}

	// Check if debt is 0 (don't waste gas)
	if repayAmount.Cmp(big.NewInt(0)) == 0 {
		log.Println("🛑 Debt is 0, skipping transaction.")
		return nil
	}

	// --- STEP 3: Execute Rescue ---

	data, err := parsedABI.Pack(
		"executeRescueWithPermit",
		common.HexToAddress(borrower),
		repayAmount, // <--- Using calculated SAFE amount
		sig.Value,   // Using original SIGNED amount for permit check
		sig.Deadline,
		sig.V,
		sig.R,
		sig.S,
	)
	if err != nil {
		return err
	}

	privateKey, err := crypto.HexToECDSA(agentPrivateKey)
	if err != nil {
		return err
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	from := crypto.PubkeyToAddress(*publicKey)

	nonce, _ := ethClient.PendingNonceAt(context.Background(), from)
	gasPrice, _ := ethClient.SuggestGasPrice(context.Background())

	tx := types.NewTransaction(
		nonce,
		agentContractAddr,
		big.NewInt(0),
		300000,
		gasPrice,
		data,
	)

	chainID, _ := ethClient.ChainID(context.Background())
	signedTx, _ := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)

	log.Println("🚀 Sending Rescue Transaction...")
	return ethClient.SendTransaction(context.Background(), signedTx)
}

/* ================= MONITOR LOOP ================= */

func monitorCrisis() {
	const agentABI = `[{"inputs":[{"internalType":"address","name":"borrower","type":"address"},{"internalType":"uint256","name":"debtToRepay","type":"uint256"},{"internalType":"uint256","name":"value","type":"uint256"},{"internalType":"uint256","name":"deadline","type":"uint256"},{"internalType":"uint8","name":"v","type":"uint8"},{"internalType":"bytes32","name":"r","type":"bytes32"},{"internalType":"bytes32","name":"s","type":"bytes32"}],"name":"executeRescueWithPermit","outputs":[],"stateMutability":"nonpayable","type":"function"}]`

	parsedABI, _ := abi.JSON(strings.NewReader(agentABI))

	query := `
    query {
      MockLending_HealthFactorUpdated(where: {newHealth: {_eq: "0"}}) {
        user
      }
    }`

	for {
		req := graphql.NewRequest(query)
		var resp struct {
			Events []CrisisUser `json:"MockLending_HealthFactorUpdated"`
		}

		if err := gqlClient.Run(context.Background(), req, &resp); err == nil {
			for _, e := range resp.Events {
				storeMutex.RLock()
				sig, ok := SignatureStore[e.User]
				storeMutex.RUnlock()

				if ok {
					log.Println("🔥 Rescuing", e.User)

					if err := executeRescue(parsedABI, e.User, sig); err != nil {
						log.Println("❌ Rescue failed:", err)
						continue
					}

					// Remove signature so we don't spam the network
					storeMutex.Lock()
					delete(SignatureStore, e.User)
					storeMutex.Unlock()

					log.Println("✅ Rescue Transaction Sent & Signature invalidated for", e.User)
				}
			}
		}
		time.Sleep(10 * time.Second) // Check every 10 seconds
	}
}

/* ================= MAIN ================= */

func main() {
	godotenv.Load()

	rpc := os.Getenv("BASE_SEPOLIA_RPC")
	ethClient, _ = ethclient.Dial(rpc)

	gqlClient = graphql.NewClient(os.Getenv("GRAPHQL_ENDPOINT"))

	agentContractAddr = common.HexToAddress(os.Getenv("AGENT_CONTRACT_ADDR"))
	agentPrivateKey = os.Getenv("PRIVATE_KEY")

	go monitorCrisis()

	http.HandleFunc("/submit-signature", enableCors(signatureHandler))

	log.Println("Agent running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}