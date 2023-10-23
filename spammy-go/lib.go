package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
)

func currentBlock(NODE_URL string) string {
	resp, err := httpGet(fmt.Sprintf("%s/block", NODE_URL))
	if err != nil {
		log.Fatalf("Failed to get current block: %v", err)
	}
	var blockRes BlockResult
	json.Unmarshal(resp, &blockRes)
	return blockRes.Result.Block.Header.Height
}

func mempoolSize(NODE_URL string) string {
	resp, err := httpGet(fmt.Sprintf("%s/num_unconfirmed_txs", NODE_URL))
	if err != nil {
		log.Fatalf("Failed to get mempool size: %v", err)
	}
	var mempoolRes MempoolResult
	json.Unmarshal(resp, &mempoolRes)
	return mempoolRes.Result.NTxs
}

func blockSize(height, NODE_URL string) []string {
	resp, err := httpGet(fmt.Sprintf("%s/block?height=%s", NODE_URL, height))
	if err != nil {
		log.Fatalf("Failed to get block size: %v", err)
	}
	var blockRes BlockResult
	json.Unmarshal(resp, &blockRes)
	return blockRes.Result.Block.Data.Txs
}

func getInitialSequence() int {
	resp, err := httpGet("http://127.0.0.1:1317/cosmos/auth/v1beta1/accounts/ADDRESS")
	if err != nil {
		log.Fatalf("Failed to get initial sequence: %v", err)
	}
	var accountRes AccountResult
	json.Unmarshal(resp, &accountRes)
	fmt.Println("sequence is", accountRes.Account.Sequence)
	return accountRes.Account.Sequence
}

func httpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// This function will load our nodes from nodes.toml
func loadNodes() []string {
	var config Config
	if _, err := toml.DecodeFile("nodes.toml", &config); err != nil {
		log.Fatalf("Failed to load nodes.toml: %v", err)
	}
	return config.SuccessfulNodes
}
