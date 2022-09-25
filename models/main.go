package models

type Block struct {
	BlockNumber       int64         `json:"blockNumber"`
	Timestamp         uint64        `json:"timestamp"`
	Difficulty        uint64        `json:"difficulty"`
	Hash              string        `json:"hash"`
	TransactionsCount int           `json:"transactionsCount"`
	Transactions      []Transaction `json:"transactions"`
}

type Transaction struct {
	Hash     string `json:"hash"`
	Value    uint64 `json:"value"`
	Gas      uint64 `json:"gas"`
	GasPrice uint64 `json:"gasPrice"`
	Nonce    uint64 `json:"nonce"`
	To       string `json:"to"`
	Pending  bool   `json:"pending"`
	Data     string `json:"data"`
}

/*
blockHash: "0xf34f0a161918d4f5f96f8e68841ddb584b3c89b090ad4930bed1095d66de4389",
  blockNumber: 21645656,
  contractAddress: null,
  cumulativeGasUsed: 39979,
  effectiveGasPrice: 11000000000,
  from: "0xce4c7c6e621cbdd1c2c5d5538bcd1054f51cf15f",
  gasUsed: 39979,
  logs: [{
      address: "0x25786cc5c07f3882a4f0b52d7900a824e6c0df82",
      blockHash: "0xf34f0a161918d4f5f96f8e68841ddb584b3c89b090ad4930bed1095d66de4389",
      blockNumber: 21645656,
      data: "0x0000000000000000000000000000000000000000000000000000001e0abc4710",
      logIndex: 0,
      removed: false,
      topics: ["", "0x000000000000000000000000ce4c7c6e621cbdd1c2c5d5538bcd1054f51cf15f", ""],
      transactionHash: "0x113dd9741c64ba9341f3009fe0538f7e7b12e465799f320994415eb3f659a5ac",
      transactionIndex: 0
  }],
  logsBloom: "",
  status: "0x1",
  to: "0x25786cc5c07f3882a4f0b52d7900a824e6c0df82",
  transactionHash: "0x113dd9741c64ba9341f3009fe0538f7e7b12e465799f320994415eb3f659a5ac",
  transactionIndex: 0,
  type: "0x0"

*/
type Receipt struct {
	From              string `json:"from"`
	To                string `json:"to"`
	Type              string `json:"type"`
	BlockNumber       string `json:"blockNumber"`
	TransactionIndex  int    `json:"transactionIndex"`
	ContractAddress   int    `json:"contractAddress"`
	Bloom             string `json:"bloom"`
	CumulativeGasUsed uint64 `json:"cumulativeGasUsed"`
	GasUsed           uint64 `json:"gasUsed"`
	Logs              []Log  `json:"logs"`
}

/*
 address: "0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c",
  blockHash: "0x3b250c45f6be373b756dd043fc2e4fe7242900eae2b6c24092b3000e8ff30327",
  blockNumber: 21645705,
  data: "0x00000000000000000000000000000000000000000000000001645ad640a03c56",
  logIndex: 0,
  removed: false,
  topics: ["", "", ""],
  transactionHash: "0xf26ec1af89e6ad702ec3076ebe52db50cc84fdbc36fea3d37858bbf058d6725b",
  transactionIndex: 0

*/
type Log struct {
	Address         string   `json:"address"`
	BlockNumber     uint64   `json:"blockNumber"`
	BlockHash       string   `json:"blockHash"`
	TransactionHash string   `json:"transactionHash"`
	Index           uint     `json:"index"`
	Data            []uint64 `json:"data"`
}
type Message struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Gas       uint64 `json:"gas"`
	GasPrice  uint64 `json:"gasPrice"`
	Value     uint64 `json:"value"`
	GasFeeCap uint64 `json:"gasFeeCap"`
	IsFake    bool   `json:"isFake"`
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

type Alert struct {
	From  string
	To    string
	Token string
	Cost  uint64
}

type Method struct {
	name     string
	methodID string
}

type Data struct {
	value uint64
}
