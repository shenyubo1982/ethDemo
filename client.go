package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {

	//client, err := ethclient.Dial("ADD_YOUR_ETHEREUM_NODE_URL")
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	fmt.Println(client)
	if err != nil {
		log.Fatalf("Oops! There was a problem", err)
	} else {
		fmt.Println("Success! you are connected to the Ethereum Network")
	}
}
