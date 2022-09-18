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

	for _, d := range latestBlock.Transactions {

		fmt.Println("From " + d.From)
		fmt.Println("To " + d.To)
		fmt.Print("Gas ")
		fmt.Println(d.Gas)
		fmt.Print("Gas Price ")
		fmt.Println(d.GasPrice)
		fmt.Print("Value ")
		fmt.Println(d.Value)
		fmt.Print("Nonce ")
		fmt.Println(d.Nonce)
		fmt.Println()
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
