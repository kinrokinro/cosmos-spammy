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
	BATCH_SIZE = 500
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

			for {
				lastBlock := currentBlock(nodeURL)
				lastBlockSize := blockSize(lastBlock, nodeURL)
				currentMempoolSize := mempoolSize(nodeURL)

				fmt.Printf("Last block height: %s, size: %d transactions\n", lastBlock, len(lastBlockSize))
				fmt.Printf("Current mempool size: %s transactions\n", currentMempoolSize)

				for i := 0; i < BATCH_SIZE; i++ {
					broadcastLog, err := sendIBCTransferViaRPC("test", nodeURL, uint64(sequence))
					if err != nil {
						log.Fatalf("Failed to send IBC transfer via RPC: %v", err)
					}
					fmt.Print(broadcastLog)

					if strings.Contains(string(broadcastLog), "code: 20") {
						fmt.Println("\033[31mMEMPOOL FULL!!!!!!!!!\033[0m")
						time.Sleep(60 * time.Second)
						break
					}

					match, _ := regexp.MatchString("account sequence mismatch", string(broadcastLog))
					if match {
						re := regexp.MustCompile(`expected (\d+)`)
						matches := re.FindStringSubmatch(string(broadcastLog))
						if len(matches) > 1 {
							sequence, err = strconv.Atoi(matches[1])
							if err != nil {
								log.Fatalf("Failed to convert sequence to integer: %v", err)
							}
							fmt.Printf("we had an account sequence mismatch, adjusting to %s\n", sequence)
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
