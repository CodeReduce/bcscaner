package ethscaner

type result struct {
	IsError        string `json:"isError"`
	ErrDescription string `json:"errDescription"`
}

type status struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Result  *result `json:"result,omitempty"`
}

type JsonRpc struct {
	JsonRpcVersion string      `json:"jsonrpc"`
	Id             uint        `json:"id"`
	Result         interface{} `json:"result"`
}

type TransactionInfo struct {
	BlockHash   string `json:"blockHash,omitempty"`
	BlockNumber string `json:"blockNumber,omitempty"`
	From        string `json:"from"`
	Gas         string `json:"gas"`
	Hash        string `json:"hash"`
	To          string `json:"to"`
}
