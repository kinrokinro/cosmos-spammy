package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
)

func currentBlock(nodeURL string) string {
	resp, err := httpGet(fmt.Sprintf("%s/block", nodeURL))
	if err != nil {
		log.Fatalf("Failed to get current block: %v", err)
	}
	var blockRes BlockResult
	err = json.Unmarshal(resp, &blockRes)
	if err != nil {
		log.Fatalf("Failed to unmarshal block result: %v", err)
	}
	return blockRes.Result.Block.Header.Height
}

func mempoolSize(nodeURL string) Result {
	resp, err := httpGet(fmt.Sprintf("%s/num_unconfirmed_txs", nodeURL))
	if err != nil {
		log.Fatalf("Failed to get mempool size: %v", err)
	}
	var mempoolRes MempoolResult
	err = json.Unmarshal(resp, &mempoolRes)
	if err != nil {
		log.Fatalf("Failed to unmarshal mempool result: %v", err)
	}
	return mempoolRes.Result
}

func blockSize(height, nodeURL string) []string {
	resp, err := httpGet(fmt.Sprintf("%s/block?height=%s", nodeURL, height))
	if err != nil {
		log.Fatalf("Failed to get block size: %v", err)
	}
	var blockRes BlockResult
	err = json.Unmarshal(resp, &blockRes)
	if err != nil {
		log.Fatalf("Failed to unmarshal block result: %v", err)
	}
	return blockRes.Result.Block.Data.Txs
}

func getInitialSequence() int {
	resp, err := httpGet("http://127.0.0.1:1317/cosmos/auth/v1beta1/accounts/ADDRESS")
	if err != nil {
		log.Fatalf("Failed to get initial sequence: %v", err)
	}
	var accountRes AccountResult
	err = json.Unmarshal(resp, &accountRes)
	if err != nil {
		log.Fatalf("Failed to unmarshal account result: %v", err)
	}
	fmt.Println("sequence is", accountRes.Account.Sequence)
	return accountRes.Account.Sequence
}

func httpGet(url string) ([]byte, error) {
	resp, err := http.Get(url) //nolint:gosec // this is what it thinks it is
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
