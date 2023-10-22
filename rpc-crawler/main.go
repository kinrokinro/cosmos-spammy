package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	initialNode     = "http://localhost:26657"
	statusEndpoint  = "/status"
	netInfoEndpoint = "/net_info"
	tomlFilename    = "successful_nodes.toml"
)

var visited = struct {
	sync.RWMutex
	nodes map[string]bool
}{nodes: make(map[string]bool)}

var successfulNodes = struct {
	sync.RWMutex
	nodes []string
}{}

var client = &http.Client{Timeout: 500 * time.Millisecond}

func main() {
	checkNode(initialNode)
	writeSuccessfulNodesToToml()
}

func checkNode(nodeAddr string) {
	if isNodeVisited(nodeAddr) {
		return
	}

	markNodeAsVisited(nodeAddr)

	resp, err := fetchNetInfo(nodeAddr)
	if err == nil {
		fmt.Println("Got net_info from", nodeAddr)
		addSuccessfulNode(nodeAddr)
	} else {
		return
	}

	for _, peer := range resp.Result.Peers {
		processPeer(&peer)
	}
}

func addSuccessfulNode(nodeAddr string) {
	successfulNodes.Lock()
	defer successfulNodes.Unlock()
	successfulNodes.nodes = append(successfulNodes.nodes, nodeAddr)
}

func normalizeAddressWithRemoteIP(nodeAddr string, remoteIP string) string {
	nodeAddr = strings.Replace(nodeAddr, "0.0.0.0", remoteIP, -1)
	nodeAddr = strings.Replace(nodeAddr, "127.0.0.1", remoteIP, -1)
	return nodeAddr
}

func isNodeVisited(nodeAddr string) bool {
	visited.RLock()
	defer visited.RUnlock()
	_, ok := visited.nodes[nodeAddr]
	return ok
}

func markNodeAsVisited(nodeAddr string) {
	visited.Lock()
	defer visited.Unlock()
	visited.nodes[nodeAddr] = true
}

func httpGet(url string) (*http.Response, error) {
	return client.Get(url)
}

func fetchNetInfo(nodeAddr string) (*NetInfoResponse, error) {
	url := nodeAddr + netInfoEndpoint
	resp, err := httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var netInfo NetInfoResponse
	err = json.NewDecoder(resp.Body).Decode(&netInfo)
	return &netInfo, err
}

func processPeer(peer *Peer) {
	rpcAddr := buildRPCAddress(peer)
	rpcAddr = normalizeAddressWithRemoteIP(rpcAddr, peer.RemoteIP)
	checkNode("http://" + rpcAddr)
}

func buildRPCAddress(peer *Peer) string {
	rpcAddr := peer.NodeInfo.Other.RPCAddress
	if strings.HasPrefix(rpcAddr, "tcp://") {
		rpcAddr = strings.TrimPrefix(rpcAddr, "tcp://")
	}
	if rpcAddr[:9] == "0.0.0.0:" || rpcAddr[:9] == "127.0.0.1:" {
		rpcAddr = peer.RemoteIP + rpcAddr[8:]
	}
	return rpcAddr
}

func getIPAddressWithoutPort(addr string) string {
	parts := strings.Split(addr, ":")
	return strings.Join(parts[:len(parts)-1], ":")
}

func writeSuccessfulNodesToToml() {
	file, err := os.Create(tomlFilename)
	if err != nil {
		fmt.Println("Error creating .toml file:", err)
		return
	}
	defer file.Close()

	successfulNodes.RLock()
	defer successfulNodes.RUnlock()

	for _, node := range successfulNodes.nodes {
		file.WriteString(fmt.Sprintf("node = \"%s\"\n", node))
	}
	fmt.Println(".toml file created with successful nodes.")
}
