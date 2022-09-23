package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
)

// 0x10ED43C718714eb63d5aA57B78B54704E256024E panckake swap router v2
// wei = 1000000000000000000
func main() {
	client, _ := ethclient.Dial("http://localhost:8545")
	memPool, _ := client.SyncProgress(context.Background())

	fmt.Println(memPool)

	os.Exit(1)
}
