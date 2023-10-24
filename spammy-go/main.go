package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"sync"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
)

const (
	BatchSize  = 100
	MaxWorkers = 1000
)

func main() {
	// tracking vars
	var successfulTxns int
	var failedTxns int
	var mu sync.Mutex
	// Declare a map to hold response codes and their counts
	responseCodes := make(map[int]int)

	// keyring
	// read seed phrase
	mnemonic, _ := os.ReadFile("seedphrase")
	// Create an in-mempory keyring

	kr := keyring.NewInMemory(nil)

	// Derive a new account from the mnemonic.
	info, err := kr.NewAccount("test", string(mnemonic), "", hd.CreateHDPath(118, 0, 0).String(), hd.Secp256k1)
	if err != nil {
		log.Fatalf("Failed to create account: %v", err)
	}

	address := info.GetAddress().String()

	sequence := getInitialSequence(address)

	successfulNodes := loadNodes()
	fmt.Printf("Number of nodes: %d\n", len(successfulNodes))

	var wg sync.WaitGroup

	// Compile the regex outside the loop
	reMismatch := regexp.MustCompile("account sequence mismatch")
	reExpected := regexp.MustCompile(`expected (\d+)`)

	for _, nodeURL := range successfulNodes {
		wg.Add(1)
		go func(nodeURL string) {
			defer wg.Done()

			currentMempoolSize := mempoolSize(nodeURL)
			fmt.Printf("Node: %s, Mempool size: %s bytes, Number of transactions: %s\n", nodeURL, currentMempoolSize.TotalBytes, currentMempoolSize.NTxs)

			startBlock := currentBlock(nodeURL)
			fmt.Printf("Script starting at block height: %s\n", startBlock)

			for {
				lastBlock := startBlock
				lastBlockSize := blockSize(lastBlock, nodeURL)
				currentMempoolSize := mempoolSize(nodeURL)

				fmt.Println("Last block height: ", lastBlock)
				fmt.Println("Last Block Size: ", lastBlockSize)
				fmt.Println(nodeURL, "Current mempool txns: %s transactions\n", currentMempoolSize.NTxs)
				fmt.Println(nodeURL, "mempool byte size:", currentMempoolSize.TotalBytes)

				var wgBatch sync.WaitGroup
				wgBatch.Add(BatchSize)

				for i := 0; i < BatchSize; i++ {
					go func() {
						defer wgBatch.Done()

						resp, _, err := sendIBCTransferViaRPC("test", nodeURL, uint64(sequence), kr)
						if err != nil {
							mu.Lock()
							failedTxns++
							mu.Unlock()
						} else {
							mu.Lock()
							successfulTxns++
							mu.Unlock()
							if resp != nil {
								// Increment the count for this response code
								mu.Lock()
								responseCodes[resp.BroadcastResult.Code]++
								mu.Unlock()
							}

							match := reMismatch.MatchString(resp.BroadcastResult.Log)
							if match {
								matches := reExpected.FindStringSubmatch(resp.BroadcastResult.Log)
								if len(matches) > 1 {
									sequence, err = strconv.Atoi(matches[1])
									if err != nil {
										log.Fatalf("Failed to convert sequence to integer: %v", err)
									}
									fmt.Printf("we had an account sequence mismatch, adjusting to %d\n", sequence)
								}
							} else {
								seqNum := sequence
								sequence = seqNum + 1
							}
						}
					}()
				}

				wgBatch.Wait()
				fmt.Println("successful transactions: ", successfulTxns)
				fmt.Println("failed transactions: ", failedTxns)
				totalTxns := successfulTxns + failedTxns
				fmt.Println("Response code breakdown:")
				for code, count := range responseCodes {
					percentage := float64(count) / float64(totalTxns) * 100
					fmt.Printf("Code %d: %d (%.2f%%)\n", code, count, percentage)
				}

				for {
					if currentBlock(nodeURL) > lastBlock {
						break
					}
				}
			}
		}(nodeURL)
	}

	wg.Wait()
}
