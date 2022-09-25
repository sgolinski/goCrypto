package handler

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"myApp/consts"
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
			fmt.Println(tx.Value().String())
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
			fmt.Println(msg.Value().String())
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
				fmt.Println("END TRANSACTION ")
			}
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
			fmt.Print("GAS ")
			fmt.Println(tx.Gas())
			fmt.Print("GAS PRICE ")
			fmt.Println(tx.GasPrice().String())
			fmt.Print("GAS FEE CAP ")
			fmt.Println(tx.GasFeeCap().String())
			fmt.Print("VALUE ")
			fmt.Println(tx.Value().String())
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
			fmt.Println(msg.Value().String())
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
				fmt.Println("END TRANSACTION ")
			}
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
			fmt.Print("GAS ")
			fmt.Println(tx.Gas())
			fmt.Print("GAS PRICE ")
			fmt.Println(tx.GasPrice().String())
			fmt.Print("GAS FEE CAP ")
			fmt.Println(tx.GasFeeCap().String())
			fmt.Print("VALUE ")
			fmt.Println(tx.Value().String())
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
			fmt.Println(msg.Value().String())
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
				fmt.Println("END TRANSACTION ")
			}
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
