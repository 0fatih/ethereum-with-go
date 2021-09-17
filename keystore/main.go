package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func createKs() {
	ks := keystore.NewKeyStore("./", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"

	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())
}

func importKs() {
	file := "UTC--2021-09-17T07-04-45.368692685Z--b461eff0b27c3e92fd8c98ed26100cad05ba4daf"

	ks := keystore.NewKeyStore("./", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())
	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}

func main() {
	createKs()
}
