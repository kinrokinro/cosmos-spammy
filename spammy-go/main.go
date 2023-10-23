package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	BatchSize = 500
)

func main() {
	successfulNodes := loadNodes()

	var wg sync.WaitGroup

	for _, nodeURL := range successfulNodes {
		wg.Add(1)
		go func(nodeURL string) {
			defer wg.Done()

			startBlock := currentBlock(nodeURL)
			fmt.Printf("Script starting at block height: %s\n", startBlock)

			sequence := getInitialSequence()

			// Compile the regex outside of the loop
			reMismatch := regexp.MustCompile("account sequence mismatch")

			//			reOneMessage := regexp.MustCompile("must contain at least one message: invalid request")

			reExpected := regexp.MustCompile(`expected (\d+)`)

			for {
				lastBlock := currentBlock(nodeURL)
				lastBlockSize := blockSize(lastBlock, nodeURL)
				currentMempoolSize := mempoolSize(nodeURL)

				fmt.Printf("Last block height: %s, size: %d transactions\n", lastBlock, len(lastBlockSize))
				fmt.Printf("Current mempool txns: %s transactions\n", currentMempoolSize.NTxs)
				fmt.Println("mempool byte size:", currentMempoolSize.TotalBytes)

				for i := 0; i < BatchSize; i++ {
					broadcastLog, reqString, err := sendIBCTransferViaRPC("test", nodeURL, uint64(sequence))
					if err != nil {
						fmt.Println(reqString)
						log.Fatalf("Failed to send IBC transfer via RPC: %v", err)
					}
					fmt.Print(broadcastLog)

					if strings.Contains(broadcastLog, "code: 20") {
						fmt.Println("\033[31mMEMPOOL FULL!!!!!!!!!\033[0m")
						time.Sleep(60 * time.Second)
						break
					}

					match := reMismatch.MatchString(broadcastLog)
					if match {
						matches := reExpected.FindStringSubmatch(broadcastLog)
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
