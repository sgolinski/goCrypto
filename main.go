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
	number := big.NewInt(int64(blockNumber))

	block, _ := client.BlockByNumber(context.Background(), number)

	for _, txn := range block.Transactions() {
		v, r, s := txn.RawSignatureValues()
		fmt.Println(v)
		fmt.Println(r)
		fmt.Println(s)
	}
}
