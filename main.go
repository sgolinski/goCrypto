package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	client, _ := ethclient.Dial("http://localhost:8545")
	blockNumber, _ := client.BlockNumber(context.Background())
	fmt.Println(blockNumber)

}
