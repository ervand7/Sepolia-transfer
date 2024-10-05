# Ethereum Transaction Sender

This Go project demonstrates how to send an Ethereum transaction (in this case, on the Sepolia test network) using the Go-Ethereum (go-ethereum) library. It also includes a utility function to generate a new Ethereum account.

## Features
- **Send Ethereum Transactions**: Send ETH to a recipient address on the Sepolia network.
- **Generate Ethereum Address**: Generate a new Ethereum account with a public address and private key.
- **Wait for Transaction to Be Mined**: Optionally wait for the transaction to be mined and get the transaction receipt.

## Prerequisites

To run this project, ensure you have the following set up:

1. **Go Programming Language**: [Install Go](https://go.dev/dl/) if you haven't already.
2. **Infura Account**: [Create an Infura account](https://infura.io/) and get an Infura Project ID to connect to the Ethereum network.
3. **Ethereum Private Key**: You need your Ethereum account's private key to sign transactions. If you don't have one, you can use the `generateETHAddress()` function in this code to create a new one.
4. **Go-Ethereum Package**: This project uses the [go-ethereum](https://github.com/ethereum/go-ethereum) package to interact with Ethereum nodes.
5. **Sepolia Testnet ETH**: You will need Sepolia testnet ETH to send transactions. Get testnet ETH from a Sepolia faucet.

## Installation

1. Clone this repository:
   ```bash
   git clone <repository-url>
   cd <repository-folder>
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Create a `.env` file to store environment variables (Infura Project ID, Private Key, Recipient Address):
   ```bash
   touch .env
   ```

4. Add the following variables to your `.env` file:
   ```bash
   INFURA_PROJECT_ID=your_infura_project_id
   PRIVATE_KEY_HEX=your_private_key_hex
   RECIPIENT_ADDRESS=recipient_ethereum_address
   ```

## Usage

1. **Sending a Transaction**:

   To send ETH from your Ethereum account to a recipient address, run the `main()` function:

   ```bash
   go run main.go
   ```

   - The program will connect to the Sepolia Ethereum testnet via Infura.
   - It will retrieve your Ethereum account address, the recipient address, nonce, and suggest gas prices.
   - After signing and sending the transaction, it will optionally wait for the transaction to be mined and output the transaction hash and block number.

2. **Generating a New Ethereum Address**:

   To generate a new Ethereum address, uncomment the `generateETHAddress()` function call in `main.go` and run:

   ```bash
   go run main.go
   ```

   This will output a new Ethereum address and private key.

## Code Breakdown

- **`weiToEther(wei *big.Int) string`**: Converts Wei (smallest unit of Ether) to Ether for better readability.
- **`main()`**:
  - Connects to Sepolia test network.
  - Loads your private key and calculates the public address.
  - Creates a new transaction to send 0.01 ETH to the recipient address.
  - Signs the transaction and sends it.
  - Optionally waits for the transaction to be mined and outputs the receipt.
- **`bindWaitMined()`**: Helper function to wait for a transaction to be mined by periodically polling for the receipt.
- **`generateETHAddress()`**: Utility function to generate a new Ethereum address and private key.

## Environment Variables

The project uses environment variables for secure storage of sensitive information:

- `INFURA_PROJECT_ID`: Infura project ID to connect to the Sepolia test network.
- `PRIVATE_KEY_HEX`: Your Ethereum private key in HEX format.
- `RECIPIENT_ADDRESS`: The recipient's Ethereum address.

Make sure you never expose your private key or `.env` file publicly.

## Dependencies

- [go-ethereum](https://github.com/ethereum/go-ethereum) - Official Go implementation of the Ethereum protocol.
- [godotenv](https://github.com/joho/godotenv) - Go package for loading environment variables from a `.env` file.

## Notes

- **Transaction Fees**: The program uses gas prices suggested by the Ethereum client. Ensure you have enough ETH in your account to cover both the transfer amount and gas fees.
- **Network**: The program currently connects to the Sepolia testnet. Modify the Infura URL in the code to connect to another Ethereum network if needed.
