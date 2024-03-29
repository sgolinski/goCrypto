package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"myApp/handler"
)

func main() {
	//client, err := ethclient.Dial("wss://silent-serene-reel.bsc-testnet.discover.quiknode.pro/5924e8eb98190e32b10ec91104e3ca71ca126f92/")
	client, err := ethclient.Dial("ws://127.0.0.1:8546")
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
			fmt.Print("Header: ")
			fmt.Println(header.Number.String())
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			handler.ExtractBlock(client, *block)
		}
	}
}
