# Zcash Transaction Processor

## Overview
This project demonstrates integrating Zcash's `zcashd` RPC functionality into a Go application for blockchain interactions. It supports fetching blockchain information, retrieving balances, and sending transactions using Go's JSON-RPC capabilities.

## Features
- Fetch blockchain details (`getblockchaininfo`)
- Retrieve wallet balance (`getbalance`)
- Send Zcash transactions (`sendtoaddress`)

## Directory Structure
```
zcash-transaction-processor/
├── main.go              # Entry point of the application
├── rpc_client/          # Handles RPC communication with zcashd
│   └── rpc_client.go    # Contains RPC logic and transaction functionality
├── config/              # Manages application configuration
│   └── config.go        # Configuration loading logic
├── go.mod               # Go module dependencies
```

## Prerequisites
- **Go**: Version 1.17+
- **zcashd**: A running `zcashd` node with RPC enabled.
- **RPC Credentials**: Ensure you have `rpcuser`, `rpcpassword`, and RPC URL.

## Configuration
Update the configuration in `config/config.go`:
```go
return &Config{
    RPCURL:      "http://127.0.0.1:8232",
    RPCUser:     "your_rpc_user",
    RPCPassword: "your_rpc_password",
}
```

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/michalstefanow/Zcash-Transaction-Processor.git
   cd Zcash-Transaction-Processor
   ```
2. Initialize Go modules:
   ```bash
   go mod tidy
   ```

## Usage
1. Run the application:
   ```bash
   go run main.go
   ```
2. The application will:
    - Fetch and display blockchain information.
    - Retrieve wallet balance.
    - Send a test transaction (ensure the `fromAddress` and `toAddress` are configured in `main.go`).

## Example Output
```plaintext
Blockchain Info: {Chain: "main", Blocks: 1500000, ...}
Balance: 12.34567890 ZEC
Transaction sent successfully. TxID: abcd1234ef567890...
```