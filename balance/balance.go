package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/bbc25118f46944c69b6980a12a299d3f")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected")

	// Get latest balances
	account := common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Latest Balance: ", balance)

	// In a specific block
	blockNumber := big.NewInt(13236630)
	balance, err = client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Balance at 5532993:", balance)

	_ = client
}
