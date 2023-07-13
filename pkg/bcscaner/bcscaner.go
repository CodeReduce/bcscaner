package bcscaner

// blockchain scaner

type BlockchainScaner interface {
	ExecutionStatus(txHash string) (bool, error)
	GetTransactionByHash(txHash string) (TransactionInfo, error)
	ERC721Transfers(contractAddr, accountAddr, sort string, page, offset int) ([]ERC721Transfer, error)
}
