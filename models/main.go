package models

import "github.com/ethereum/go-ethereum/core/types"

type Block struct {
	BlockNumber       int64         `json:"blockNumber"`
	Timestamp         uint64        `json:"timestamp"`
	Difficulty        uint64        `json:"difficulty"`
	Hash              string        `json:"hash"`
	TransactionsCount int           `json:"transactionCount"`
	Transactions      []Transaction `json:"transactions"`
}

/*
blockHash: "0xb00f99f737d3f7ce145a8de3d14e8cdb04e79ae1f1fd3e1cddcab8e8e92400cf",
  blockNumber: 21439247,
  from: "0x2dca8b8ccfc3686a1ea1ba0c6ee11f15a0e968cb",
  gas: 21000,
  gasPrice: 10000000000,
  hash: "0xdaaad06df9d64e21391c5f3ca2b9ff4cfceea33a4da07cc3ec9cac73a50085ec",
  input: "0x",
  nonce: 155172,
  r: "0x6c9daba596b3b9261e794be4f7b86306f2810e2a6364052d3b1a9fb9d609fdb7",
  s: "0x1e361193e1d47ca1df66d208fd262f6835a541d1f0b3fab426defc28fa2a6b46",
  to: "0xebe4489af7a8707527fcbfbca9ce60b824cd1039",
  transactionIndex: 0,
  type: "0x0",
  v: "0x94",
  value: 800000000000000
*/
type Transaction struct {
	BlockNumber int64            `json:"blockNumber"`
	From        string           `json:"from"`
	Hash        string           `json:"hash"`
	Value       string           `json:"value"`
	Gas         uint64           `json:"gas"`
	GasPrice    uint64           `json:"gasPrice"`
	Input       string           `json:"input"`
	Nonce       uint64           `json:"nonce"`
	To          string           `json:"to"`
	Pending     bool             `json:"pending"`
	AccessList  types.AccessList `json:"accessList"`
}

type TransferEthRequest struct {
	PrivKey string `json:"privKey"`
	To      string `json:"to"`
	Amount  int64  `json:"amount"`
}
type HashResponse struct {
	Hash string `json:"hash"`
}

type BalanceResponse struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
	Symbol  string `json:"symbol"`
	Units   string `json:"units"`
}

type Error struct {
	Code    uint64 `json:"code"`
	Message string `json:"message"`
}
