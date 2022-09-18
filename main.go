package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func main() {

	client, _ := ethclient.Dial("http://localhost:8545")
	blockNumber, _ := client.BlockNumber(context.Background())

	block, _ := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))

	fmt.Println(block.Transactions())

}
