package main

import (
	"context"
	"crypto/ecdsa"
	
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// ---------------- CONFIG & GLOBALS ----------------
var (
	client              *ethclient.Client
	privateKey          *ecdsa.PrivateKey
	agentAddress        common.Address
	chainID             *big.Int
	
	// Contract Addresses
	debtTokenAddr       common.Address
	guardianManagerAddr common.Address
	rescueAgentAddr     common.Address

	// In-Memory Store for Signatures (User -> SignatureData)
	signatureStore = make(map[string]SignatureData)
	storeMutex     sync.RWMutex
)

// ---------------- DATA STRUCTURES ----------------

type SignatureData struct {
	Owner    string `json:"owner"`
	Value    string `json:"value"`
	Deadline string `json:"deadline"`
	V        uint8  `json:"v"`
	R        string `json:"r"`
	S        string `json:"s"`
}

// ---------------- ABIs (Minimal for A2A Flow) ----------------

// 1. ERC-20 Permit ABI
const tokenABIJson = `[
	{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"},{"internalType":"uint256","name":"deadline","type":"uint256"},{"internalType":"uint8","name":"v","type":"uint8"},{"internalType":"bytes32","name":"r","type":"bytes32"},{"internalType":"bytes32","name":"s","type":"bytes32"}],"name":"permit","outputs":[],"stateMutability":"nonpayable","type":"function"}
]`

// 2. GuardianManager ABI
const managerABIJson = `[
	{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"address","name":"user","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"},{"internalType":"address","name":"targetAgent","type":"address"}],"name":"authorizeAgent","outputs":[],"stateMutability":"nonpayable","type":"function"}
]`

// 3. RescueAgent ABI (Updated for A2A: payer + debtor)
const rescueABIJson = `[
	{"inputs":[{"internalType":"address","name":"payer","type":"address"},{"internalType":"address","name":"debtor","type":"address"},{"internalType":"uint256","name":"amountToRepay","type":"uint256"},{"internalType":"uint256","name":"agentFee","type":"uint256"},{"internalType":"uint256","name":"value","type":"uint256"},{"internalType":"uint256","name":"deadline","type":"uint256"},{"internalType":"uint8","name":"v","type":"uint8"},{"internalType":"bytes32","name":"r","type":"bytes32"},{"internalType":"bytes32","name":"s","type":"bytes32"}],"name":"executeRescueWithPermit","outputs":[],"stateMutability":"nonpayable","type":"function"}
]`

// ---------------- MAIN ENTRY POINT ----------------

func main() {
	// 1. Load Env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  Warning: No .env file found")
	}

	// 2. Connect to Blockchain
	setupBlockchain()

	// 3. Start Background Monitor
	go monitorCrisisLoop()

	// 4. Start HTTP Server for Signatures
	setupServer()
}

// ---------------- SETUP FUNCTIONS ----------------

func setupBlockchain() {
	var err error
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		log.Fatal("❌ RPC_URL is missing in .env")
	}

	client, err = ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal("❌ Failed to connect to EthClient:", err)
	}

	// Load Private Key
	pkStr := os.Getenv("PRIVATE_KEY")
	privateKey, err = crypto.HexToECDSA(strings.TrimPrefix(pkStr, "0x"))
	if err != nil {
		log.Fatal("❌ Invalid Private Key:", err)
	}

	// Derive Agent Public Address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("❌ Error casting public key")
	}
	agentAddress = crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get Chain ID
	chainID, err = client.NetworkID(context.Background())
	if err != nil {
		log.Fatal("❌ Failed to get ChainID:", err)
	}

	// Parse Contract Addresses
	debtTokenAddr = common.HexToAddress(os.Getenv("DEBT_TOKEN_ADDR"))
	guardianManagerAddr = common.HexToAddress(os.Getenv("GUARDIAN_MANAGER_ADDR"))
	rescueAgentAddr = common.HexToAddress(os.Getenv("RESCUE_AGENT_ADDR"))

	log.Printf("🤖 Agent initialized at: %s", agentAddress.Hex())
	log.Printf("🛡️  Manager: %s | 🚑 Rescue: %s", guardianManagerAddr.Hex(), rescueAgentAddr.Hex())
}

func setupServer() {
	r := gin.Default()

	// FIX CORS: Allow Flutter app to talk to Codespace
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	r.Use(cors.New(config))

	// Endpoint to receive permits
	r.POST("/submit-signature", func(c *gin.Context) {
		var sig SignatureData
		if err := c.BindJSON(&sig); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storeMutex.Lock()
		signatureStore[strings.ToLower(sig.Owner)] = sig
		storeMutex.Unlock()

		log.Printf("✅ Stored A2A signature for %s", sig.Owner)
		c.JSON(http.StatusOK, gin.H{"status": "received"})
	})

	log.Println("🚀 Server listening on :8081")
	r.Run(":8081")
}

// ---------------- MONITORING LOOP ----------------

func monitorCrisisLoop() {
	// Simple simulation loop. In production, query Envio/Graph here.
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		storeMutex.RLock()
		// Iterate over all known users we have signatures for
		for user, sig := range signatureStore {
			// TODO: Query Envio here to check real health factor.
			// For Hackathon Demo: We assume anyone with a sig is a target if we run this.
			// You can add a boolean flag or a specific API call to trigger this.
			
			log.Printf("🔍 Checking user %s...", user)
			
			// --- REAL RESCUE TRIGGER ---
			// Uncomment the 'if' condition below to connect to real indexer logic
			// if (checkUserHealth(user) < 1.0) {
				go performA2ARescue(user, sig)
			// }
		}
		storeMutex.RUnlock()
	}
}

// ---------------- A2A RESCUE SEQUENCE ----------------

func performA2ARescue(userAddrStr string, sig SignatureData) {
	log.Printf("🚨 INITIATING A2A RESCUE FOR %s", userAddrStr)

	// 1. Prepare Data
	userAddr := common.HexToAddress(userAddrStr)
	
	val, _ := new(big.Int).SetString(sig.Value, 10)
	deadline, _ := new(big.Int).SetString(sig.Deadline, 10)
	
	// R and S conversion
	rVal := common.HexToHash(sig.R)
	sVal := common.HexToHash(sig.S)

	// Get fresh nonce for the Agent
	nonce, err := client.PendingNonceAt(context.Background(), agentAddress)
	if err != nil {
		log.Printf("❌ Failed to get nonce: %v", err)
		return
	}

	// ---------------------------------------------------------
	// TX 1: SUBMIT PERMIT to DEBT TOKEN
	// (User -> Spender: GuardianManager)
	// ---------------------------------------------------------
	log.Println("➡️ Step 1: Submitting Permit to Token Contract...")
	parsedTokenABI, _ := abi.JSON(strings.NewReader(tokenABIJson))
	
	// Note: Spender in the permit must be GUARDIAN MANAGER for A2A
	input1, _ := parsedTokenABI.Pack("permit", 
		userAddr, 
		guardianManagerAddr, // <--- Spender is Manager
		val, deadline, sig.V, rVal, sVal,
	)

	tx1Hash := sendTx(debtTokenAddr, input1, nonce)
	if tx1Hash == "" { return }
	nonce++ // Increment nonce locally for next tx

	// ---------------------------------------------------------
	// TX 2: AUTHORIZE AGENT (Manager -> RescueAgent)
	// ---------------------------------------------------------
	log.Println("➡️ Step 2: Manager Authorizing RescueRunner...")
	parsedManagerABI, _ := abi.JSON(strings.NewReader(managerABIJson))
	
	// Amount to rescue (e.g., 500 USDC). For demo, we use a fixed amount or part of permit
	amountToRescue := new(big.Int).Div(val, big.NewInt(2)) // Rescue 50% of permit limit

	input2, _ := parsedManagerABI.Pack("authorizeAgent", 
		debtTokenAddr, 
		userAddr, 
		amountToRescue, 
		rescueAgentAddr, // <--- Target Agent B
	)

	tx2Hash := sendTx(guardianManagerAddr, input2, nonce)
	if tx2Hash == "" { return }
	nonce++

	// ---------------------------------------------------------
	// TX 3: EXECUTE RESCUE (RescueAgent Logic)
	// ---------------------------------------------------------
	log.Println("➡️ Step 3: Executing Final Rescue...")
	parsedRescueABI, _ := abi.JSON(strings.NewReader(rescueABIJson))
	
	agentFee := big.NewInt(0) // Free for hackathon demo
	
	// Payer = Guardian Manager
	// Debtor = User
	input3, _ := parsedRescueABI.Pack("executeRescueWithPermit",
		guardianManagerAddr, // Payer
		userAddr,            // Debtor
		amountToRescue,
		agentFee,
		val, deadline, sig.V, rVal, sVal, // Pass dummy/old permit data just to satisfy interface
	)

	tx3Hash := sendTx(rescueAgentAddr, input3, nonce)
	if tx3Hash != "" {
		log.Printf("🏆 A2A RESCUE COMPLETE! Final Tx: %s", tx3Hash)
		
		// Remove signature to prevent loops
		storeMutex.Lock()
		delete(signatureStore, strings.ToLower(sig.Owner))
		storeMutex.Unlock()
	}
}

// ---------------- HELPER: SEND TRANSACTION ----------------

func sendTx(to common.Address, data []byte, nonce uint64) string {
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	// Bump gas for Testnet reliability
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(2)) 

	// Create TX
	tx := types.NewTransaction(nonce, to, big.NewInt(0), 300000, gasPrice, data)

	// Sign TX
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Printf("❌ Signing error: %v", err)
		return ""
	}

	// Broadcast
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Printf("❌ Broadcast error: %v", err)
		return ""
	}

	log.Printf("   ✅ Tx Sent: %s", signedTx.Hash().Hex())
	
	// Wait a bit to ensure sequence order on testnet
	time.Sleep(2 * time.Second) 
	return signedTx.Hash().Hex()
}