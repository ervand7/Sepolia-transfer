package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}
}

// weiToEther converts Wei to Ether for readability
func weiToEther(wei *big.Int) string {
	ether := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(1e18))
	return ether.Text('f', 18)
}

func main() {
	// Retrieve Infura Project ID and Private Key from environment variables
	infuraProjectID := os.Getenv("INFURA_PROJECT_ID")
	privateKeyHex := os.Getenv("PRIVATE_KEY_HEX")

	// Connect to Sepolia Ethereum network via Infura
	infuraURL := fmt.Sprintf("https://sepolia.infura.io/v3/%s", infuraProjectID)
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer client.Close()
	fmt.Println("Successfully connected to Sepolia Ethereum network")

	// Load the private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}

	// Get the public key and address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Failed to cast public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("From Address: %s\n", fromAddress.Hex())

	// Specify the recipient address
	toAddress := common.HexToAddress(os.Getenv("RECIPIENT_ADDRESS"))
	fmt.Printf("To Address: %s\n", toAddress.Hex())

	// Get the nonce for the sender address
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}
	fmt.Printf("Nonce: %d\n", nonce)

	// Define the amount to send (in Wei)
	// Example: 0.01 ETH
	value := big.NewInt(10000000000000000) // 0.01 ETH in Wei
	fmt.Printf("ETH transfer amount: %s\n", weiToEther(value))

	// Suggest Gas Price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}
	fmt.Printf("gasPrice: %d\n", gasPrice)

	// Estimate Gas Limit (21000 is standard for ETH transfer)
	gasLimit := uint64(21000)

	// Create the transaction
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	// Get the chain ID for Sepolia (chain ID = 11155111)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get network ID: %v", err)
	}
	fmt.Printf("chainID: %s\n", chainID)

	// Sign the transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	// Send the transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	fmt.Printf("Transaction sent! Tx Hash: %s\n", signedTx.Hash().Hex())

	// Optional: Wait for the transaction to be mined
	receipt, err := bindWaitMined(context.Background(), client, signedTx)
	if err != nil {
		log.Fatalf("Failed to wait for transaction to be mined: %v", err)
	}

	if receipt.Status == types.ReceiptStatusFailed {
		log.Fatal("Transaction failed")
	}

	fmt.Printf("Transaction mined! Block Number: %v\n", receipt.BlockNumber)
}

func bindWaitMined(ctx context.Context, client *ethclient.Client, tx *types.Transaction) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(ctx, tx.Hash())
		if err == nil {
			return receipt, nil
		}
		fmt.Println("Waiting for transaction to be mined...")
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			time.Sleep(time.Second)
		}
	}
}

func generateETHAddress() {
	// Generation of a new private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	// Obtaining a public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Failed to obtain public key from private key")
	}

	// Generating an address from a public key
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Exporting a private key to HEX format
	privateKeyHex := fmt.Sprintf("%x", crypto.FromECDSA(privateKey))

	fmt.Println("=== New Ethereum account ===")
	fmt.Printf("Address: %s\n", address.Hex())
	fmt.Printf("Private key (HEX): %s\n", privateKeyHex)
}
