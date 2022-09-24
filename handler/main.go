package handler

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
)

// 0x10ED43C718714eb63d5aA57B78B54704E256024E panckake swap router v2
// wei = 1000000000000000000
func ExtractBlock(client *ethclient.Client, block types.Block) {

	for _, tx := range block.Transactions() {

		err, msg := createMsg(client, tx)

		fmt.Print("From ")
		fmt.Println(msg.From().String())
		fmt.Print("To ")
		fmt.Println(msg.To().String())
		fmt.Println("Txn " + tx.Hash().String())
		// cost is real value of transaction
		fmt.Print("Cost ")
		//cost := tx.Cost()
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

func createMsg(client *ethclient.Client, tx *types.Transaction) (error, types.Message) {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), nil)
	return err, msg
}
