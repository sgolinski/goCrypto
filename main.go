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

	//var transactions []types.Transaction

	for _, txn := range block.Transactions() {
		fmt.Print("Hash ")
		fmt.Println(txn.Hash().String())
		fmt.Print("Value ")
		fmt.Println(txn.Value().String())
		fmt.Print("Access list ")
		fmt.Println(txn.AccessList().StorageKeys())
		fmt.Print("Nonce ")
		fmt.Println(txn.Nonce())
		fmt.Print("To ")
		fmt.Println(txn.To().String())
		fmt.Print("Gas price ")
		fmt.Println(txn.GasPrice().String())
		fmt.Print("Type ")
		fmt.Println(txn.Type())
		fmt.Print("Chain ID ")
		fmt.Println(txn.ChainId().String())
		fmt.Print("Cost ")
		fmt.Println(txn.Cost().String())
		fmt.Print("Data")
		fmt.Println(txn.UnmarshalBinary(txn.Data()))

		fmt.Println()
	}

}
