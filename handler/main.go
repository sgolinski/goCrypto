package handler

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/status-im/keycard-go/hexutils"
	"log"
	"myApp/consts"
	"strconv"
	"strings"
)

func ExtractBlock(client *ethclient.Client, block types.Block) {

	for _, tx := range block.Transactions() {

		value := tx.Value().Uint64()

		//if value > 10000000000000000000 {
		txData := common.Bytes2Hex(tx.Data())

		if consts.ContainsRemoveLiquidityMethod(txData) {
			err, msg := createMsg(client, tx)
			fmt.Println(msg.From().String())
			fmt.Println(msg.To().String())
			fmt.Println(msg.Value().Uint64())
			if err != nil {
				log.Fatal(err)
			}
			receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				log.Fatal(err)
			}
			for _, lg := range receipt.Logs {
				fmt.Print("DATA ")
				data := hexutils.BytesToHex(lg.Data)
				sliceData := extractHexSlice(data)
				fmt.Println(sliceData)

			}
			fmt.Println("END TRANSACTION ")
		} else if consts.ContainsAddLiquidityMethod(txData) {
			err, msg := createMsg(client, tx)
			fmt.Println(msg.From().String())
			fmt.Println(msg.To().String())
			fmt.Println(msg.Value().Uint64())
			if err != nil {
				log.Fatal(err)
			}
			receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				log.Fatal(err)
			}
			for _, lg := range receipt.Logs {
				data := hexutils.BytesToHex(lg.Data)
				sliceData := extractHexSlice(data)
				fmt.Println(sliceData)

			}
			fmt.Println("END TRANSACTION ")
			fmt.Println()
		} else if consts.ContainsSwapMethod(txData) && value > 10000000000000000000 {
			err, msg := createMsg(client, tx)
			fmt.Println(msg.From().String())
			fmt.Println(msg.To().String())
			fmt.Println(msg.Value().Uint64())
			if err != nil {
				log.Fatal(err)
			}
			receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				log.Fatal(err)
			}
			for _, lg := range receipt.Logs {
				fmt.Print("DATA ")
				data := hexutils.BytesToHex(lg.Data)
				sliceData := extractHexSlice(data)
				fmt.Println(sliceData)

			}
			fmt.Println("END TRANSACTION ")
		}

	}

}

func createMsg(client *ethclient.Client, tx *types.Transaction) (error, types.Message) {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), nil)
	return err, msg
}

func hexaNumberToInteger(hexaString string) string {

	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	return numberStr
}

func convertHexToDecimal(hexaNumber string) uint64 {

	output, err := strconv.ParseUint(hexaNumberToInteger(hexaNumber), 16, 64)
	if err != nil {
		fmt.Println(err)
	}
	return output
}

func extractHexSlice(hexaString string) []uint64 {

	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	numberStr = strings.ToLower(numberStr)
	counter := len(numberStr)
	divide := counter / 64

	slice := make([]uint64, divide)
	start := 0
	end := 64
	for i := 0; i < divide; i++ {
		s := numberStr[start:end]
		start += 64
		end += 64
		val := convertHexToDecimal(s)
		slice[i] = val
	}
	return slice
}
