package polyscaner

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/CodeReduce/bcscaner/pkg/bcscaner"
)

type Polyscan struct {
	host   string
	apiKey string
}

const (
	moduleAccount              = "account"
	moduleTransaction          = "transaction"
	moduleProxy                = "proxy"
	actionGetStatus            = "gettxreceiptstatus"
	actionGetTokenList         = "tokennfttx"
	actionGetTransactionByHash = "eth_getTransactionByHash"
)

func NewRepo(hostApi, apiKey string) bcscaner.BlockchainScaner {
	return &Polyscan{
		host:   hostApi,
		apiKey: apiKey,
	}
}

func (e *Polyscan) ExecutionStatus(txHash string) (bool, error) {
	u := url.URL{
		Scheme: "https",
		Host:   e.host,
		Path:   "api",
	}
	rq := u.Query()
	rq.Set("module", moduleTransaction)
	rq.Set("action", actionGetStatus)
	rq.Set("txhash", txHash)
	rq.Set("apikey", e.apiKey)
	u.RawQuery = rq.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return false, fmt.Errorf("execution status method response with status code: %v, message: %v", res.StatusCode, res.Status)
	}
	var s status
	err = json.NewDecoder(res.Body).Decode(&s)
	if err != nil {
		return false, fmt.Errorf("decode json error: %v", err)
	}

	if s.Result == nil {
		return false, fmt.Errorf("pending. result is nil")
	}
	if s.Message != messageOK {
		return false, nil
	}
	return s.Result.Status == "1", nil
}

func (e *Polyscan) GetTransactionByHash(txHash string) (bcscaner.TransactionInfo, error) {
	u := url.URL{
		Scheme: "https",
		Host:   e.host,
		Path:   "api",
	}
	rq := u.Query()
	rq.Set("module", moduleProxy)
	rq.Set("action", actionGetTransactionByHash)
	rq.Set("txhash", txHash)
	rq.Set("apikey", e.apiKey)
	u.RawQuery = rq.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		return bcscaner.TransactionInfo{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return bcscaner.TransactionInfo{}, fmt.Errorf("execution status method response with status code: %v, message: %v", res.StatusCode, res.Status)
	}
	var j bcscaner.JsonRpc
	err = json.NewDecoder(res.Body).Decode(&j)
	if err != nil {
		return bcscaner.TransactionInfo{}, fmt.Errorf("decode json error: %v", err)
	}

	txInfoData, err := json.Marshal(j.Result)
	if err != nil {
		return bcscaner.TransactionInfo{}, fmt.Errorf("marshal json error: %v", err)
	}

	var txinfo bcscaner.TransactionInfo
	err = json.Unmarshal(txInfoData, &txinfo)
	if err != nil {
		return bcscaner.TransactionInfo{}, fmt.Errorf("unmarshal json error: %v", err)
	}
	return txinfo, nil
}

func (e *Polyscan) ERC721Transfers(contractAddr, accountAddr, sort string, page, offset int) ([]bcscaner.ERC721Transfer, error) {
	u := url.URL{
		Scheme: "https",
		Host:   e.host,
		Path:   "api",
	}
	rq := u.Query()
	rq.Set("module", moduleAccount)
	rq.Set("action", actionGetTokenList)
	rq.Set("contractaddress", contractAddr)
	rq.Set("address", accountAddr)
	rq.Set("startblock", "0")
	rq.Set("endblock", "99999999")
	rq.Set("page", fmt.Sprint(page))
	rq.Set("offset", fmt.Sprint(offset))
	rq.Set("sort", sort)
	rq.Set("apikey", e.apiKey)
	u.RawQuery = rq.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("error making http request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("execution status method response with status code: %v, message: %v", res.StatusCode, res.Status)
	}
	var j bcscaner.JsonRpc
	err = json.NewDecoder(res.Body).Decode(&j)
	if err != nil {
		return nil, fmt.Errorf("decode json error: %v", err)
	}

	erc721TransfersData, err := json.Marshal(j.Result)
	if err != nil {
		return nil, fmt.Errorf("marshal json error: %v", err)
	}

	erc721Transfers := []bcscaner.ERC721Transfer{}
	err = json.Unmarshal(erc721TransfersData, &erc721Transfers)
	if err != nil {
		return nil, fmt.Errorf("unmarshal json error: %v", err)
	}
	return erc721Transfers, nil
}
