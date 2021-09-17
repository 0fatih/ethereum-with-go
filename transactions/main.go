package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_PROJECT_ID")
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(1345675)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	// Transaction information
	for _, tx := range block.Transactions() {
		fmt.Println("Hash: ", tx.Hash().Hex())
		fmt.Println("Value: ", tx.Value().String())
		fmt.Println("Gas: ", tx.Gas())
		fmt.Println("Gas Price: ", tx.GasPrice().Uint64())
		fmt.Println("Nonce: ", tx.Nonce())
		fmt.Println("Data: ", tx.Data())
		fmt.Println("To: ", tx.To().Hex())
		if err != nil {
			log.Fatal(err)
		}

		if msg, err := tx.AsMessage(types.NewEIP155Signer(tx.ChainId()), big.NewInt(1)); err == nil {
			fmt.Println("From: ", msg.From().Hex())
		}

		// receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println("Receipt: ", receipt.Status)

		fmt.Println()
	}

	// Transaction count at specific block hash
	blockHash := common.HexToHash("0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	// Just transaction hashes
	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tx.Hash().Hex())
	}

	// Is pending
	txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tx.Hash().Hex())
	fmt.Println(isPending)

}
