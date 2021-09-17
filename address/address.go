package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func main() {

	address := common.HexToAddress("0x6598bbadC1EffF11e4cd3E5C14CC530e30F7F2F0")
	fmt.Println(address.Hex())
	fmt.Println(address.Hash().Hex())
	fmt.Println(address.Bytes())
}
