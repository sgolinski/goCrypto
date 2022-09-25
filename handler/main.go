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
		data := common.Bytes2Hex(tx.Data())

		if consts.ContainsRemoveLiquidityMethod(data) {
			fmt.Println("TRANSACTION REMOVE LIQUIDUTY ")
			fmt.Print("HASH ")
			fmt.Println(tx.Hash().String())
			fmt.Print("COST ")
			fmt.Println(tx.Cost().String())
			fmt.Print("NONCE ")
			fmt.Println(tx.Nonce())
			fmt.Print("TO ")
			fmt.Println(tx.To().String())
			fmt.Print("GAS ")
			fmt.Println(tx.Gas())
			fmt.Print("GAS PRICE ")
			fmt.Println(tx.GasPrice().String())
			fmt.Print("GAS FEE CAP ")
			fmt.Println(tx.GasFeeCap().String())
			fmt.Print("VALUE ")
			fmt.Println(tx.Value().Uint64())
			fmt.Print("DATA ")
			fmt.Println(common.BytesToHash(tx.Data()))
			fmt.Print("TYPE ")
			fmt.Println(tx.Type())

			err, msg := createMsg(client, tx)

			fmt.Println("Message: ")
			fmt.Print("TO ")
			fmt.Println(msg.To().String())
			fmt.Print("FROM ")
			fmt.Println(msg.From().String())
			fmt.Print("GAS ")
			fmt.Println(msg.Gas())
			fmt.Print("VALUE ")
			fmt.Println(msg.Value().Uint64())
			fmt.Print("IS FAKE ")
			fmt.Println(msg.IsFake())
			fmt.Print("GAS FEE CAP ")
			fmt.Println(msg.GasFeeCap().String())
			fmt.Print("GAS PRICE ")
			fmt.Println(msg.GasPrice().String())
			fmt.Print("NONCE ")
			fmt.Println(msg.Nonce())
			fmt.Println("END MESSAGE")

			if err != nil {
				log.Fatal(err)
			}
			receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				log.Fatal(err)
			}
			for _, lg := range receipt.Logs {
				fmt.Println()
				fmt.Println("RECEIPT ")
				fmt.Print("Address ")
				fmt.Println(lg.Address)
				address := lg.Address
				fmt.Print("BLOCK NUMBER ")
				fmt.Println(lg.BlockNumber)
				fmt.Print("BLOCK HASH ")
				fmt.Println(lg.BlockHash.String())
				fmt.Print("TXHASH ")
				fmt.Println(lg.TxHash.String())
				contract, _ := client.BalanceAt(context.Background(), address, nil)
				fmt.Print("CONTRACT: ")
				fmt.Println(contract.String())
				fmt.Print("DATA ")
				data := hexutils.BytesToHex(tx.Data())
				sliceData := extractHexSlice(data)
				fmt.Println(sliceData)

			}
			fmt.Println("END TRANSACTION ")
			fmt.Println()
		} else if consts.ContainsAddLiquidityMethod(data) {
			fmt.Println("TRANSACTION  ADD LIQUIDUTY ")
			fmt.Print("HASH ")
			fmt.Println(tx.Hash().String())
			fmt.Print("COST ")
			fmt.Println(tx.Cost().String())
			fmt.Print("NONCE ")
			fmt.Println(tx.Nonce())
			fmt.Print("TO ")
			fmt.Println(tx.To().String())
			fmt.Print("VALUE ")
			fmt.Println(tx.Value().Uint64())
			fmt.Print("DATA ")
			fmt.Println(common.BytesToHash(tx.Data()))

			err, msg := createMsg(client, tx)

			fmt.Println("Message: ")
			fmt.Print("TO ")
			fmt.Println(msg.To().String())
			fmt.Print("FROM ")
			fmt.Println(msg.From().String())
			fmt.Print("VALUE ")
			fmt.Println(msg.Value().String())
			fmt.Print("IS FAKE ")
			fmt.Println(msg.IsFake())
			fmt.Println("END MESSAGE")

			if err != nil {
				log.Fatal(err)
			}
			receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				log.Fatal(err)
			}
			for _, lg := range receipt.Logs {
				fmt.Println()
				fmt.Println("RECEIPT ")
				fmt.Print("Address ")
				fmt.Println(lg.Address)
				address := lg.Address
				fmt.Print("TXHASH ")
				fmt.Println(lg.TxHash.String())
				fmt.Print("Contract ")
				fmt.Println(lg.Address)
				contract, _ := client.BalanceAt(context.Background(), address, nil)
				fmt.Print("CONTRACT: ")
				fmt.Println(contract.String())
				fmt.Print("DATA ")
				data := hexutils.BytesToHex(tx.Data())
				sliceData := extractHexSlice(data)
				fmt.Println(sliceData)

			}
			fmt.Println("END TRANSACTION ")
			fmt.Println()
		} else if consts.ContainsSwapMethod(data) && value > 10000000000000000000 {
			fmt.Println("TRANSACTION SWAP ")
			fmt.Print("HASH ")
			fmt.Println(tx.Hash().String())
			fmt.Print("COST ")
			fmt.Println(tx.Cost().String())
			fmt.Print("NONCE ")
			fmt.Println(tx.Nonce())
			fmt.Print("TO ")
			fmt.Println(tx.To().String())
			fmt.Print("VALUE ")
			fmt.Println(tx.Value().String())
			fmt.Print("DATA ")
			fmt.Println(common.BytesToHash(tx.Data()))
			err, msg := createMsg(client, tx)

			fmt.Println("Message: ")
			fmt.Print("TO ")
			fmt.Println(msg.To().String())
			fmt.Print("FROM ")
			fmt.Println(msg.From().String())
			fmt.Print("VALUE ")
			fmt.Println(msg.Value().String())
			fmt.Print("IS FAKE ")
			fmt.Println(msg.IsFake())

			if err != nil {
				log.Fatal(err)
			}
			receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				log.Fatal(err)
			}
			for _, lg := range receipt.Logs {
				fmt.Println()
				fmt.Println("RECEIPT ")
				fmt.Print("Address ")
				fmt.Println(lg.Address)
				address := lg.Address
				fmt.Print("BLOCK NUMBER ")
				fmt.Println(lg.BlockNumber)
				fmt.Print("BLOCK HASH ")
				fmt.Println(lg.BlockHash.String())
				fmt.Print("TOPICS ")
				fmt.Println(lg.Topics)
				fmt.Print("REMOVED ")
				fmt.Println(lg.Removed)
				fmt.Print("INDEX ")
				fmt.Println(lg.Index)
				fmt.Print("TXHASH ")
				fmt.Println(lg.TxHash.String())
				fmt.Print("TXINDEX ")
				fmt.Println(lg.TxIndex)
				fmt.Print("Contract ")
				fmt.Println(lg.Address)
				contract, _ := client.BalanceAt(context.Background(), address, nil)
				fmt.Print("CONTRACT: ")
				fmt.Println(contract)
				fmt.Print("DATA ")
				data := hexutils.BytesToHex(tx.Data())
				sliceData := extractHexSlice(data)
				fmt.Println(sliceData)

			}
			fmt.Println("END TRANSACTION ")
			fmt.Println()
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
