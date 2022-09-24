package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"myapp/handler"
)

// 0x10ED43C718714eb63d5aA57B78B54704E256024E panckake swap router v2
// wei = 1000000000000000000
func main() {
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
			fmt.Println(header)
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			handler.ExtractBlock(client, *block)
		}
	}
}
