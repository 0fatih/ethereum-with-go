package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/YOUR_PROJECT_ID")
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("[+]new block added:")
			fmt.Println("block hash: ", block.Hash().Hex())
			fmt.Println("block number: ", block.Number().Uint64())
			fmt.Println("block time: ", block.Time())
			fmt.Println("block nonce: ", block.Nonce())
			fmt.Println("transaction count: ", len(block.Transactions()))
			fmt.Println()
		}
	}
}
