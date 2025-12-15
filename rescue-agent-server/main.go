package main

import (
	"context"
	"encoding/hex"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"

    "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/machinebox/graphql" 
)

// --- GLOBAL STATE & CONFIG ---
var (
	// In-memory store for user signatures received from the Flutter Web dApp
	SignatureStore = make(map[string]SignatureData)
	storeMutex     sync.RWMutex

	// Clients
	ethClient *ethclient.Client
	gqlClient *graphql.Client
	
	// Configuration
	agentContractAddr common.Address
	agentPrivateKey   string
)

// SignatureData struct maps directly to the required fields from the Permit message
type SignatureData struct {
	UserAddress string `json:"owner"`
	Value       *big.Int `json:"value"` // Max amount authorized
	Deadline    *big.Int `json:"deadline"`
	V           uint8    `json:"v"`
	R           [32]byte `json:"r"`
	S           [32]byte `json:"s"`
}

// CrisisUser struct maps directly to the GraphQL response fields
type CrisisUser struct {
	User        string `json:"user"`
	BlockNumber string `json:"block_number"`
}

// --- HTTP HANDLER: RECEIVING SIGNATURES FROM FLUTTER WEB ---
// --- NEW CORS MIDDLEWARE FUNCTION ---
func enableCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	}
}


func signatureHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		SignatureData
		RHex string `json:"r"` // Receive R as hex string
		SHex string `json:"s"` // Receive S as hex string
	}
	
	// Decode JSON data from Flutter POST request
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, fmt.Sprintf("Invalid JSON: %v", err), http.StatusBadRequest)
		return
	}

	// Convert hex strings back to byte arrays for Ethereum
	rBytes, errR := hex.DecodeString(strings.TrimPrefix(data.RHex, "0x"))
	sBytes, errS := hex.DecodeString(strings.TrimPrefix(data.SHex, "0x"))
	if errR != nil || errS != nil || len(rBytes) != 32 || len(sBytes) != 32 {
		http.Error(w, "Invalid R or S format. Must be 32-byte hex string.", http.StatusBadRequest)
		return
	}

	// Save the signature data
	storeMutex.Lock()
	defer storeMutex.Unlock()
	
	copy(data.R[:], rBytes)
	copy(data.S[:], sBytes)
	
	SignatureStore[data.UserAddress] = data.SignatureData
	log.Printf("Successfully stored permit signature for user: %s", data.UserAddress)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Signature stored successfully."))
}


// --- CORE LOGIC: GRAPHQL MONITORING AND EXECUTION ---

func monitorCrisisLoop() {
	// ABI of the executeRescueWithPermit function (extracted from your RescueAgent contract)
	const agentABI = `[{"inputs":[{"internalType":"address","name":"borrower","type":"address"},{"internalType":"uint256","name":"debtToRepay","type":"uint256"},{"internalType":"uint256","name":"value","type":"uint256"},{"internalType":"uint256","name":"deadline","type":"uint256"},{"internalType":"uint8","name":"v","type":"uint8"},{"internalType":"bytes32","name":"r","type":"bytes32"},{"internalType":"bytes32","name":"s","type":"bytes32"}],"name":"executeRescueWithPermit","outputs":[],"stateMutability":"nonpayable","type":"function"}]`

	parsedABI, err := abi.JSON(strings.NewReader(agentABI))
	if err != nil {
		log.Fatalf("Failed to parse agent ABI: %v", err)
	}

	// GraphQL query to find users in crisis (e.g., health factor is 0)
	const crisisQuery = `
		query CurrentCrisisUsers {
			MockLending_HealthFactorUpdated(
				where: {newHealth: {_eq: "0"}}
			) {
				user
			}
		}`

	for {
		req := graphql.NewRequest(crisisQuery)
		var response struct {
			CrisisEvents []CrisisUser `json:"MockLending_HealthFactorUpdated"`
		}

		if err := gqlClient.Run(context.Background(), req, &response); err != nil {
			log.Printf("GraphQL Query Error: %v", err)
			time.Sleep(30 * time.Second)
			continue
		}

		for _, event := range response.CrisisEvents {
			user := common.HexToAddress(event.User).Hex()
			
			storeMutex.RLock()
			sigData, exists := SignatureStore[user]
			storeMutex.RUnlock()

			if exists {
				log.Printf("🔥 CRISIS DETECTED for user %s. Initiating rescue...", user)
				if err := executeRescueTx(parsedABI, user, sigData); err != nil {
					log.Printf("🚨 Rescue FAILED for %s: %v", user, err)
				} else {
					log.Printf("✅ Rescue SUCCESS for %s. Removing signature.", user)
					storeMutex.Lock()
					delete(SignatureStore, user) // Prevent re-use of signature
					storeMutex.Unlock()
				}
			}
		}

		time.Sleep(60 * time.Second) // Poll every 60 seconds
	}
}

// --- CORE LOGIC: EXECUTE TRANSACTION ---

func executeRescueTx(parsedABI abi.ABI, borrower string, sigData SignatureData) error {
	// Repay amount (Example: 100 tokens, replace with actual required amount logic)
	debtToRepay := big.NewInt(0)
	debtToRepay.SetString("100000000000000000000", 10) 

	// 1. Prepare Calldata: ABI encoding of the function and arguments
	packedData, err := parsedABI.Pack(
		"executeRescueWithPermit",
		common.HexToAddress(borrower),
		debtToRepay,
		sigData.Value,     // The max authorized amount
		sigData.Deadline,
		sigData.V,
		sigData.R,
		sigData.S,
	)
	if err != nil {
		return fmt.Errorf("failed to pack data: %w", err)
	}

	// 2. Load Agent Wallet and Prepare Transaction Metadata
	// 2. Load Agent Wallet and Prepare Transaction Metadata
	privateKey, err := crypto.HexToECDSA(agentPrivateKey)
	if err != nil {
		return fmt.Errorf("invalid private key in environment: %w", err)
	}
	
	// FIX: Directly get the public key and then its address.
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		return fmt.Errorf("cannot assert public key to type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress) // Get transaction nonce
	if err != nil {
		return fmt.Errorf("failed to get nonce: %w", err)
	}
	gasPrice, err := ethClient.SuggestGasPrice(context.Background()) // Get gas price
	if err != nil {
		return fmt.Errorf("failed to get gas price: %w", err)
	}

	// 3. Create, Sign, and Send Transaction
	tx := types.NewTransaction(
		nonce,
		agentContractAddr,
		big.NewInt(0), // Value is 0 (no native token sent)
		300000,        // Gas Limit (example)
		gasPrice,
		packedData,
	)

	chainID, err := ethClient.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %w", err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey) // Sign the transaction
	if err != nil {
		return fmt.Errorf("failed to sign transaction: %w", err)
	}

	if err := ethClient.SendTransaction(context.Background(), signedTx); err != nil { // Send the raw transaction
		return fmt.Errorf("failed to send transaction: %w", err)
	}

	log.Printf("Transaction sent! Hash: %s", signedTx.Hash().Hex())
	return nil
}

// --- MAIN SETUP ---

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		// Log fatal if .env is missing or invalid
		log.Fatal("Error loading .env file. Ensure it is in the project root and configured.")
	}
	
	// Initialize Clients and Config
	var err error
	rpcUrl := os.Getenv("BASE_SEPOLIA_RPC")
	if rpcUrl == "" {
		log.Fatal("BASE_SEPOLIA_RPC is not set in .env")
	}
	ethClient, err = ethclient.Dial(rpcUrl) // Connect to RPC
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}
	
	gqlEndpoint := os.Getenv("GRAPHQL_ENDPOINT")
	if gqlEndpoint == "" {
		log.Fatal("GRAPHQL_ENDPOINT is not set in .env")
	}
	gqlClient = graphql.NewClient(gqlEndpoint) // Connect to GraphQL indexer
	
	agentContractAddr = common.HexToAddress(os.Getenv("AGENT_CONTRACT_ADDR"))
	agentPrivateKey = os.Getenv("PRIVATE_KEY")
	if agentPrivateKey == "" {
		log.Fatal("PRIVATE_KEY is not set in .env")
	}

	// Start the background monitoring worker
	go monitorCrisisLoop()

	// Start the HTTP server to receive signatures from Flutter Web
	http.HandleFunc("/submit-signature", enableCors(signatureHandler))
	
	log.Println("Agent Server running on http://localhost:8081. Monitoring crisis events...")
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", nil))
}