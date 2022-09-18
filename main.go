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

		fmt.Println("Message")
		fmt.Print("From ")
		fmt.Println(msg.From().String())
		fmt.Print("To ")
		fmt.Println(msg.To().String())
		fmt.Println("Txn " + tx.Hash().String())
		fmt.Print(tx.Gas() / 1000000000000000000)
		fmt.Println(" bnb")
		if err != nil {
			log.Fatal(err)
		}

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		for _, lg := range receipt.Logs {
			fmt.Println()
			fmt.Print("Address ")
			fmt.Println(lg.Address)
			fmt.Print("BlockNumber ")
			fmt.Println(lg.BlockNumber)
			fmt.Print("BlockHash ")
			fmt.Println(lg.BlockHash.String())
			fmt.Print("Index ")
			fmt.Println(lg.Index)
			fmt.Print("Removed ")
			fmt.Println(lg.Removed)
			fmt.Println()
		}

		fmt.Println()

	}
	os.Exit(1)
}
