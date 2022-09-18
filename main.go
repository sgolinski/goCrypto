package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"log"
	Handlers "main/handler"
	"main/modules"
	"net/http"
)

func main() {
	// Create a client instance to connect to our providr
	client, err := ethclient.Dial("http://localhost:8545")

	if err != nil {
		fmt.Println(err)
	}

	latestBlock := modules.GetLatestBlock(*client)
	transactions := latestBlock.Transactions

	for _, t := range transactions {

		fmt.Println(t.Hash)
		fmt.Println(t.Value)
		fmt.Println(t.From)
		fmt.Println(t.To)
		fmt.Println(t.Gas)
		fmt.Println(t.GasPrice)
		fmt.Println(t.Input)
		fmt.Println(t.Nonce)
		fmt.Println(t.Pending)

	}

	if err != nil {
		fmt.Println(err)
	}

	// Create a mux router
	r := mux.NewRouter()

	// We will define a single endpoint
	r.Handle("/api/v1/eth/{module}", Handlers.ClientHandler{client})
	log.Fatal(http.ListenAndServe(":8080", r))

}
