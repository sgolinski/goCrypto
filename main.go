package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
)

// 0x10ED43C718714eb63d5aA57B78B54704E256024E panckake swap router v2
// wei = 1000000000000000000
func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	memPool, _ := client.SyncProgress(context.Background())

	fmt.Println(memPool)
	if err != nil {
		log.Fatal(err)
	}

	b, _ := client.BlockNumber(context.Background())
	blockNumber := big.NewInt(int64(b))
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), nil)

		fmt.Print("From ")
		fmt.Println(msg.From().String())
		fmt.Print("To ")
		fmt.Println(msg.To().String())
		fmt.Println("Txn " + tx.Hash().String())

		// cost is real value of transaction
		fmt.Print("Cost ")
		cost := tx.Cost().Uint64()

		fmt.Print(cost / 1000000000000000000)
		fmt.Println(" gwei")
		if err != nil {
			log.Fatal(err)
		}

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		for _, lg := range receipt.Logs {
			fmt.Print("Address ")
			fmt.Println(lg.Address)
			address := lg.Address
			contract, _ := client.BalanceAt(context.Background(), address, nil)
			fmt.Println(contract)
		}

		fmt.Println()

	}
	os.Exit(1)
}
