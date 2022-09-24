package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

// 0x10ED43C718714eb63d5aA57B78B54704E256024E panckake swap router v2
// wei = 1000000000000000000
func main() {
	client, err := ethclient.Dial("ws://127.0.0.1:8546")
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	// PObranie balansu konta
	//blockNumber, _ := client.BlockNumber(context.Background())
	//bigIntBlockNumber := big.NewInt(int64(blockNumber))
	//address := common.HexToAddress("0xf5dFe65e388694B5928e1f82f6A644Be63bD702a")
	//balanceAt, _ := client.BalanceAt(context.Background(), address, bigIntBlockNumber)

	//eth syncing
	//sync, _ := client.SyncProgress(context.Background())
	//fmt.Println(sync)

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println()
			fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Time())              // 1529525947
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7
			fmt.Println()
		}
	}
}
