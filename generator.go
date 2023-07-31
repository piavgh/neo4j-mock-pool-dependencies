package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateAddress() string {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("Address:", address)

	return address
}

func GeneratePoolType() string {
	rand.Seed(time.Now().UnixNano())

	in := []string{"curve-meta", "curve-oracle", "gmx"}
	randomIndex := rand.Intn(len(in))
	pick := in[randomIndex]

	return pick
}
