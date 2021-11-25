package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os/exec"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Settings struct {
	AddrBridge_RMRK string `json:"AddrBridge_RMRK"`
	AddrHub         string `json:"AddrHub"`
	PrivateKey      string `json:"PrivateKey"`
	RPC             string `json:"RPC"`
}

var Setting Settings

func main() {
	//Upload settings
	content, err := ioutil.ReadFile("Settings.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	err = json.Unmarshal(content, &Setting)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	log.Println(">Settings loaded")

	client, err := ethclient.Dial(Setting.RPC)
	if err != nil {
		log.Fatal("Error connecting: ", err)
	}

	privateKey, err := crypto.HexToECDSA(Setting.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(Setting.AddrHub)
	main_hub, err := NewHub(address, client)
	if err != nil {
		log.Fatal(err)
	}

	for {
		tx, err := main_hub.GetMOVROrders(&bind.CallOpts{})
		if err != nil {
			log.Fatal(err)
		}

		var count uint = 0
		for i, s := range tx {
			log.Println(i, s.UID, s.Address)
			cmd := exec.Command("/usr/bin/npx", "-y", "ts-node", "./run-simple-script.ts", s.UID, s.Address)
			var buf bytes.Buffer
			cmd.Stdout = &buf
			err := cmd.Start()
			if err != nil {
				fmt.Printf("error: %v\n", err)
			}
			err = cmd.Wait()
			log.Println(buf.String())
			count = count + 1
		}

		if count > 0 {
			nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
			if err != nil {
				log.Fatal(err)
			}

			gasPrice, err := client.SuggestGasPrice(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			auth := bind.NewKeyedTransactor(privateKey)
			auth.Nonce = big.NewInt(int64(nonce))
			auth.Value = big.NewInt(0)     // in wei
			auth.GasLimit = uint64(300000) // in units
			auth.GasPrice = gasPrice

			tx2, err := main_hub.FillMOVROrders(auth, big.NewInt(int64(count)))
			if err != nil {
				log.Fatal(err)
			}

			log.Println("Close orders:", tx2.Hash().Hex())
		}
		time.Sleep(5 * time.Minute)
	}
}
