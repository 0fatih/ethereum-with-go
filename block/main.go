package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/cf0e3418fb49483ea007fa174d2f553e")
	if err != nil {
		log.Fatal(err)
	}

	// Block Header
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Latest block header: ", header.Number.String())

	// Full Block
	blockNumber := big.NewInt(13241939)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Block Number: ", block.Number().Uint64(), "\nBlock Time: ", block.Time(), "\nBlock Difficulty: ", block.Difficulty().Uint64(), "\nBlock Hash: ", block.Hash().Hex(), "\nTransactions: ", len(block.Transactions()))
}
