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
	client, err := ethclient.Dial("http://localhost:7545")

	if err != nil {
		fmt.Println(err)
	}

	transaction := modules.GetLatestBlock(*client)
	fmt.Println(transaction)

	balance, _ := modules.GetAddressBalance(*client, address)
	fmt.Println(balance)

	if err != nil {
		fmt.Println(err)
	}

	// Create a mux router
	r := mux.NewRouter()

	// We will define a single endpoint
	r.Handle("/api/v1/eth/{module}", Handlers.ClientHandler{client})
	log.Fatal(http.ListenAndServe(":8080", r))

}
