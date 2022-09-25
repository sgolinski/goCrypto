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
	"myApp/models"
	"strconv"
	"strings"
)

func ExtractBlock(client *ethclient.Client, block types.Block) {

	for _, tx := range block.Transactions() {

		value := tx.Value().Uint64()

		//if value > 10000000000000000000 {
		txData := common.Bytes2Hex(tx.Data())

		if consts.ContainsRemoveLiquidityMethod(txData) {
			fmt.Println("TRANSACTION REMOVE LIQUIDUTY ")

			//hash := tx.Hash().String()
			//nonce := tx.Nonce()
			//to := tx.To().String()
			//gas := tx.Gas()
			//gasPrice := tx.GasPrice().Uint64()
			//txValue := tx.Value().Uint64()

			//transaction := models.Transaction{
			//	Hash:     hash,
			//	Value:    txValue,
			//	Gas:      gas,
			//	GasPrice: gasPrice,
			//	Nonce:    nonce,
			//	To:       to,
			//	Pending:  false,
			//	Data:     txData,
			//}
			err, msg := createMsg(client, tx)

			msgTo := msg.To().String()
			msgFrom := msg.From().String()
			msgGas := msg.Gas()
			msgValue := msg.Value().Uint64()
			msgIsFake := msg.IsFake()
			msgFeeCap := msg.GasFeeCap().Uint64()
			msgGasPrice := msg.GasPrice().Uint64()

			message := models.Message{
				From:      msgFrom,
				To:        msgTo,
				Gas:       msgGas,
				GasPrice:  msgGasPrice,
				Value:     msgValue,
				GasFeeCap: msgFeeCap,
				IsFake:    msgIsFake,
			}

			if err != nil {
				log.Fatal(err)
			}
			receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				log.Fatal(err)
			}
			count := len(receipt.Logs)

			logs := make([]models.Log, count)

			for i, lg := range receipt.Logs {
				logAddress := lg.Address
				logBlockNumber := lg.BlockNumber
				logBlockHash := lg.BlockHash.String()
				logTransactionHash := lg.TxHash.String()

				data := hexutils.BytesToHex(lg.Data)
				sliceData := extractHexSlice(data)
				l := models.Log{
					Address:         logAddress.String(),
					BlockNumber:     logBlockNumber,
					BlockHash:       logBlockHash,
					Index:           lg.Index,
					TransactionHash: logTransactionHash,
					Data:            sliceData,
				}
				logs[i] = l
			}

			fmt.Printf("Transaction from %s to %s value %d", message.From, message.To, message.Value)
		} else if consts.ContainsAddLiquidityMethod(txData) {
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
			fmt.Println(tx.Value())
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
				data := hexutils.BytesToHex(lg.Data)
				sliceData := extractHexSlice(data)
				fmt.Println(sliceData)

			}
			fmt.Println("END TRANSACTION ")
			fmt.Println()
		} else if consts.ContainsSwapMethod(txData) && value > 10000000000000000000 {
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
				data := hexutils.BytesToHex(lg.Data)
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
