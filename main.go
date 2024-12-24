package main

import (
	"fmt"
	"log"

	"zcash-transaction-processor/config"
	"zcash-transaction-processor/rpc_client"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize the RPC client
	client, err := rpc_client.NewRPCClient(cfg.RPCURL, cfg.RPCUser, cfg.RPCPassword)
	if err != nil {
		log.Fatalf("Failed to initialize RPC client: %v", err)
	}

	// Fetch and display blockchain info
	info, err := client.GetBlockchainInfo()
	if err != nil {
		log.Fatalf("Failed to fetch blockchain info: %v", err)
	}
	fmt.Printf("Blockchain Info: %+v\n", info)

	// Example: Fetch balance
	balance, err := client.GetBalance()
	if err != nil {
		log.Fatalf("Failed to fetch balance: %v", err)
	}
	fmt.Printf("Balance: %.8f ZEC\n", balance)

	// Example: Send transaction
	txID, err := client.SendTransaction("fromAddress", "toAddress", 0.1)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}
	fmt.Printf("Transaction sent successfully. TxID: %s\n", txID)
}
