// Package dataCollection provides a suite of functions to access the official INFURA API
package dataCollection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/LucaPaterlini/infura/config"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// apiCallPOST call the third party api with a timeout, and returns the content of the http response.
func apiCallPOST(jsonStr []byte, requestTimeout time.Duration) (statusCode int, header map[string][]string, body []byte, err error) {
	url := config.FullMainNetPath

	client := &http.Client{Timeout: requestTimeout}

	var resp *http.Response
	resp, err = client.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return
	}
	body, _ = ioutil.ReadAll(resp.Body)

	_ = resp.Body.Close()

	statusCode = resp.StatusCode
	header = resp.Header
	return
}

// GetBlock using the third party api gets the data of the requested block,
// if the third party is not able to provide an answer within the requested timeout it returns timeout error.
func GetBlock(blockNumber uint64, requestTimeout time.Duration) (int, map[string][]string, []byte,  error) {
	var jsonStr = []byte(fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getBlockByNumber",
		"params": ["0x%x",false],"id":1}`, blockNumber))
	return apiCallPOST(jsonStr, requestTimeout)
}

// GetTransaction using the third party api gets the data of the requested transaction,
// if the third party is not able to provide an answer within the requested timeout it returns timeout error.
func GetTransaction(blockNumber, index uint64, requestTimeout time.Duration) (int, map[string][]string, []byte,  error) {
	var jsonStr = []byte(fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getTransactionByBlockNumberAndIndex",
		"params": ["0x%x","0x0"],"id":%d}`, blockNumber, index))
	return apiCallPOST(jsonStr, requestTimeout)
}

// GetLastBlockNumber using the third party api gets the last block,
// if the third party is not able to provide an answer within the requested timeout it returns timeout error.
func GetLastBlockNumber(requestTimeout time.Duration) (lastBlock uint64, err error) {
	var jsonStr = []byte(fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_blockNumber","params": [],"id":1}`))
	var body []byte
	_, _, body, err = apiCallPOST(jsonStr, requestTimeout)
	if err != nil {
		return
	}
	type Response struct {
		Result string
	}
	var resp Response
	_ = json.Unmarshal(body, &resp)
	lastBlock, err = strconv.ParseUint(resp.Result[2:], 16, 64)
	return
}
