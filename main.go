package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func main() {

	client, _ := ethclient.Dial("http://localhost:8545")

	blockNumber, _ := client.BlockNumber(context.Background())
	number := big.NewInt(int64(blockNumber))

	block, _ := client.BlockByNumber(context.Background(), number)

	for _, txn := range block.Transactions() {

		hash := txn.Hash().String()
		fmt.Println("Hash " + hash)

		value := txn.Value().String()
		fmt.Println("Value " + value)

		v, r, s := txn.RawSignatureValues()
		fmt.Println(hexutil.EncodeBig(v))
		fmt.Println(hexutil.EncodeBig(r))
		fmt.Println(hexutil.EncodeBig(s))

	}
}
