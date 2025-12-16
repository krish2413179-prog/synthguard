package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
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
	 rescuedUsers = make(map[string]bool)
	signatureStore = make(map[string]SignatureData)
	storeMutex     sync.RWMutex

	ethClient *ethclient.Client
	gqlClient *graphql.Client

	agentContractAddr common.Address
	agentPrivateKey   string

	lendingContractAddr = common.HexToAddress("0xa0f95A73BA2c1395E9F4B95e6F6b7faF3E07A447")
)

/* ================= DATA STRUCTS ================= */

type SignatureData struct {
	Owner    common.Address
	Value    *big.Int
	Deadline *big.Int
	V        uint8
	R        [32]byte
	S        [32]byte
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

/* ================= SIGNATURE INGEST ================= */

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

	value, ok := new(big.Int).SetString(payload.Value, 10)
	if !ok {
		http.Error(w, "invalid value", http.StatusBadRequest)
		return
	}

	deadline, ok := new(big.Int).SetString(payload.Deadline, 10)
	if !ok {
		http.Error(w, "invalid deadline", http.StatusBadRequest)
		return
	}

	if time.Now().Unix() > deadline.Int64() {
		http.Error(w, "signature expired", http.StatusBadRequest)
		return
	}

	v := payload.V
	if v < 27 {
		v += 27
	}

	rBytes, _ := hex.DecodeString(strings.TrimPrefix(payload.R, "0x"))
	sBytes, _ := hex.DecodeString(strings.TrimPrefix(payload.S, "0x"))

	var rArr, sArr [32]byte
	copy(rArr[:], rBytes)
	copy(sArr[:], sBytes)

	owner := common.HexToAddress(payload.Owner)

	storeMutex.Lock()
	signatureStore[strings.ToLower(owner.Hex())] = SignatureData{
		Owner:    owner,
		Value:    value,
		Deadline: deadline,
		V:        v,
		R:        rArr,
		S:        sArr,
	}
	storeMutex.Unlock()

	log.Println("✅ Stored signature for", owner.Hex())
	w.WriteHeader(http.StatusOK)
}

/* ================= RESCUE EXECUTION ================= */

func executeRescue(parsedABI abi.ABI, borrower common.Address, sig SignatureData) error {
	if borrower != sig.Owner {
		return errors.New("borrower != signature owner")
	}

	if time.Now().Unix() > sig.Deadline.Int64() {
		return errors.New("permit expired")
	}

	const lendingABIJSON = `
	[
	  {
		"inputs":[{"internalType":"address","name":"user","type":"address"}],
		"name":"debt",
		"outputs":[{"internalType":"uint256","name":"","type":"uint256"}],
		"stateMutability":"view",
		"type":"function"
	  }
	]`

	lendingABI, _ := abi.JSON(strings.NewReader(lendingABIJSON))
	dataDebt, _ := lendingABI.Pack("debt", borrower)

	result, err := ethClient.CallContract(context.Background(), ethereum.CallMsg{
		To:   &lendingContractAddr,
		Data: dataDebt,
	}, nil)
	if err != nil {
		return err
	}

	out, _ := lendingABI.Unpack("debt", result)
	actualDebt := out[0].(*big.Int)

	if actualDebt.Sign() == 0 {
		log.Println("🛑 No debt, skipping rescue")
		return nil
	}

	// Debt to repay (never exceeds signed permit value)
	repayAmount := new(big.Int).Set(actualDebt)
	if repayAmount.Cmp(sig.Value) > 0 {
		repayAmount = sig.Value
	}

	// 💰 Agent fee (set to zero for now, or configure)
	agentFee := new(big.Int).Div(repayAmount, big.NewInt(200))

	// Safety check: permit must cover debt + fee
	total := new(big.Int).Add(repayAmount, agentFee)
	if sig.Value.Cmp(total) < 0 {
		return errors.New("permit value < debt + agent fee")
	}

	log.Println(
		"💸 Repaying:", repayAmount.String(),
		"| Fee:", agentFee.String(),
		"| Permit:", sig.Value.String(),
	)

	// 🔥 CORRECT ABI ENCODING (8 PARAMS)
	data, err := parsedABI.Pack(
		"executeRescueWithPermit",
		borrower,
		repayAmount,
		agentFee,   // ✅ REQUIRED PARAM
		sig.Value,  // 🔒 EXACT SIGNED VALUE
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

	log.Println("🚀 Sending rescue tx for", borrower.Hex())
	return ethClient.SendTransaction(context.Background(), signedTx)
}
func getDebt(user common.Address) (*big.Int, error) {
	const lendingABIJSON = `
	[
	  {
	    "inputs":[{"internalType":"address","name":"user","type":"address"}],
	    "name":"debt",
	    "outputs":[{"internalType":"uint256","name":"","type":"uint256"}],
	    "stateMutability":"view",
	    "type":"function"
	  }
	]`

	lendingABI, err := abi.JSON(strings.NewReader(lendingABIJSON))
	if err != nil {
		return nil, err
	}

	data, _ := lendingABI.Pack("debt", user)

	result, err := ethClient.CallContract(context.Background(), ethereum.CallMsg{
		To:   &lendingContractAddr,
		Data: data,
	}, nil)
	if err != nil {
		return nil, err
	}

	out, err := lendingABI.Unpack("debt", result)
	if err != nil {
		return nil, err
	}

	return out[0].(*big.Int), nil
}

/* ================= MONITOR ================= */

func monitorCrisis() {
	const agentABIJSON = `
	[
	  {
	    "inputs": [
	      { "internalType": "address", "name": "user", "type": "address" },
	      { "internalType": "uint256", "name": "amountToRepay", "type": "uint256" },
	      { "internalType": "uint256", "name": "agentFee", "type": "uint256" },
	      { "internalType": "uint256", "name": "value", "type": "uint256" },
	      { "internalType": "uint256", "name": "deadline", "type": "uint256" },
	      { "internalType": "uint8", "name": "v", "type": "uint8" },
	      { "internalType": "bytes32", "name": "r", "type": "bytes32" },
	      { "internalType": "bytes32", "name": "s", "type": "bytes32" }
	    ],
	    "name": "executeRescueWithPermit",
	    "outputs": [],
	    "stateMutability": "nonpayable",
	    "type": "function"
	  }
	]`

	parsedABI, err := abi.JSON(strings.NewReader(agentABIJSON))
	if err != nil {
		log.Fatal("❌ Failed to parse Agent ABI:", err)
	}

	// ⚠️ IMPORTANT:
	// Do NOT filter here, indexer returns historical rows
	query := `
	query {
	  MockLending_HealthFactorUpdated {
	    user
	  }
	}`

	log.Println("🧠 Crisis monitor started")

	for {
		req := graphql.NewRequest(query)

		var resp struct {
			Events []CrisisUser `json:"MockLending_HealthFactorUpdated"`
		}

		err := gqlClient.Run(context.Background(), req, &resp)
		if err != nil {
			log.Println("❌ GraphQL error:", err)
			time.Sleep(10 * time.Second)
			continue
		}

		for _, e := range resp.Events {
			userAddr := common.HexToAddress(e.User)
			key := strings.ToLower(userAddr.Hex())
			var debt *big.Int
	debt, err = getDebt(userAddr)
	if err != nil {
		log.Println("❌ Failed to fetch debt:", err)
		continue
	}

			log.Println("🧪 Crisis event detected for:", key)

			// ------------------------------------------------
			// 1️⃣ DEDUPLICATION (CRITICAL)
			// ------------------------------------------------
debt, err = getDebt(userAddr)
if err != nil {
	log.Println("❌ Failed to fetch debt:", err)
	continue
}

if debt.Sign() == 0 {
	// Crisis resolved → allow future rescues
	if rescuedUsers[key] {
		log.Println("🔄 Debt cleared, resetting rescue flag for", key)
		delete(rescuedUsers, key)
	}
	continue
}
			
			if rescuedUsers[key] {
				log.Println("⏭ Already rescued, skipping:", key)
				continue
			}

			// ------------------------------------------------
			// 2️⃣ CHECK AUTHORIZATION
			// ------------------------------------------------
			storeMutex.RLock()
			sig, ok := signatureStore[key]
			storeMutex.RUnlock()

			if !ok {
				log.Println("⛔ No authorization stored for:", key)
				continue
			}

			// ------------------------------------------------
			// 3️⃣ PERMIT EXPIRY CHECK
			// ------------------------------------------------
			if time.Now().Unix() > sig.Deadline.Int64() {
				log.Println("⏰ Permit expired for:", key)
				storeMutex.Lock()
				delete(signatureStore, key)
				storeMutex.Unlock()
				continue
			}

			// ------------------------------------------------
			// 4️⃣ EXECUTE RESCUE
			// ------------------------------------------------
			log.Println("🔥 Executing rescue for:", key)

			err := executeRescue(parsedABI, userAddr, sig)
			if err != nil {
				log.Println("❌ Rescue failed for", key, ":", err)
				continue
			}

			// ------------------------------------------------
			// 5️⃣ MARK AS HANDLED
			// ------------------------------------------------
			rescuedUsers[key] = true
			log.Println("✅ Rescue completed for:", key)
		}

		time.Sleep(10 * time.Second)
	}
}





			

		

		
/* ================= MAIN ================= */

func main() {
	godotenv.Load()

	ethClient, _ = ethclient.Dial(os.Getenv("BASE_SEPOLIA_RPC"))
	gqlClient = graphql.NewClient(os.Getenv("GRAPHQL_ENDPOINT"))

	agentContractAddr = common.HexToAddress(os.Getenv("AGENT_CONTRACT_ADDR"))
	agentPrivateKey = os.Getenv("PRIVATE_KEY")

	go monitorCrisis()

	http.HandleFunc("/submit-signature", enableCors(signatureHandler))

	log.Println("🤖 Agent running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
