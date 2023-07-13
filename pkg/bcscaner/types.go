package bcscaner

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

// ERC721Transfer holds info from ERC721 token transfer event query
type ERC721Transfer struct {
	BlockNumber       int    `json:"blockNumber,string"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             int    `json:"nonce,string"`
	BlockHash         string `json:"blockHash"`
	From              string `json:"from"`
	ContractAddress   string `json:"contractAddress"`
	To                string `json:"to"`
	TokenID           string `json:"tokenID"`
	TokenName         string `json:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol"`
	TokenDecimal      uint8  `json:"tokenDecimal,string"`
	TransactionIndex  int    `json:"transactionIndex,string"`
	Gas               int    `json:"gas,string"`
	GasPrice          string `json:"gasPrice"`
	GasUsed           int    `json:"gasUsed,string"`
	CumulativeGasUsed int    `json:"cumulativeGasUsed,string"`
	Input             string `json:"input"`
	Confirmations     int    `json:"confirmations,string"`
}
