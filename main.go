package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

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

		fmt.Println(tx.Hash().String())

		v, r, s := tx.RawSignatureValues()

		value := common.BigToAddress(v)
		receicer := common.BigToAddress(r)
		sender := common.BigToAddress(s)
		fmt.Print("value: ")
		fmt.Println(value)
		fmt.Print("Reciver: ")
		fmt.Println(receicer)
		fmt.Print("Sender: ")
		fmt.Println(sender)
		fmt.Println(msg.From().String())
		fmt.Println(msg.To().String())
		fmt.Println(msg.Gas())
		fmt.Println(msg.GasPrice().String())
		fmt.Println(msg.Value().String())
		fmt.Println(msg.GasFeeCap().String())
		fmt.Println(msg.IsFake())
		fmt.Println(msg.AccessList().StorageKeys())
		fmt.Println()
		if err != nil {
			log.Fatal(err)
		}

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(receipt.Status) // 1
	}

	//blockHash := common.HexToHash("0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9")
	//count, err := client.TransactionCount(context.Background(), blockHash)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for idx := uint(0); idx < count; idx++ {
	//	tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	//}
	//
	//txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")
	//tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	//fmt.Println(isPending)       // false
}
