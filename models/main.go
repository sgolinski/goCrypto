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
	Value    string `json:"value"`
	Gas      uint64 `json:"gas"`
	GasPrice uint64 `json:"gasPrice"`
	Nonce    uint64 `json:"nonce"`
	To       string `json:"to"`
	Pending  bool   `json:"pending"`
	Data     string `json:"data"`
}
type Log struct {
	Address     string `json:"address"`
	BlockNumber int64  `json:"blockNumber"`
	BlockHash   string `json:"blockHash"`
	Index       int    `json:"index"`
	Removed     int    `json:"removed"`
}

type Receipt struct {
	Status            string `json:"status"`
	Type              string `json:"type"`
	BlockNumber       string `json:"blockNumber"`
	TransactionIndex  int    `json:"transactionIndex"`
	ContractAddress   int    `json:"contractAddress"`
	Bloom             string `json:"bloom"`
	CumulativeGasUsed uint64 `json:"cumulativeGasUsed"`
	GasUsed           uint64 `json:"gasUsed"`
	Logs              []Log  `json:"logs"`
}

type Message struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Gas       uint64 `json:"gas"`
	GasPrice  uint64 `json:"gasPrice"`
	Value     uint64 `json:"value"`
	GasFeeCap string `json:"gasFeeCap"`
	IsFake    int    `json:"isFake"`
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
