package main

import (
	"fmt"
	"os"
	"time"

	"github.com/0fatih/ethereum-with-go/erc20Token/fetch"
)

func main() {
	start := time.Now()
	if len(os.Args) < 2 {
		fmt.Println("Enter an address")
		os.Exit(1)
	}
	fetch.FetchBalances(os.Args[1])
	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("Finished in %s", elapsed))
}
