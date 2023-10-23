package main

type Header struct {
	Height string `json:"height"`
}

type Data struct {
	Txs []string `json:"txs"`
}

type Block struct {
	Header Header `json:"header"`
	Data   Data   `json:"data"`
}

type ResultBlock struct {
	Block Block `json:"block"`
}

type BlockResult struct {
	Result ResultBlock `json:"result"`
}

type MempoolResult struct {
	Result Result `json:"result"`
}

type BroadcastRequest struct {
	Jsonrpc                string `json:"jsonrpc"`
	ID                     string `json:"id"`
	Method                 string `json:"method"`
	BroadcastRequestParams `json:"params"`
}

type BroadcastRequestParams struct {
	Tx string `json:"tx"`
}

type BroadcastResponse struct {
	Jsonrpc         string `json:"jsonrpc"`
	ID              int    `json:"id"`
	BroadcastResult `json:"result"`
}

type BroadcastResult struct {
	Code      int    `json:"code"`
	Data      string `json:"data"`
	Log       string `json:"log"`
	Codespace string `json:"codespace"`
	Hash      string `json:"hash"`
}

type Result struct {
	NTxs       string `json:"n_txs"`
	TotalBytes string `json:"total_bytes"`
}

type Account struct {
	Sequence int `json:"sequence"`
}

type AccountResult struct {
	Account Account `json:"account"`
}
type Transaction struct {
	Body       Body     `json:"body"`
	AuthInfo   AuthInfo `json:"auth_info"`
	Signatures []string `json:"signatures"`
}

type Body struct {
	Messages                    []Message `json:"messages"`
	Memo                        string    `json:"memo"`
	TimeoutHeight               string    `json:"timeout_height"`
	ExtensionOptions            []string  `json:"extension_options"`
	NonCriticalExtensionOptions []string  `json:"non_critical_extension_options"`
}

type Message struct {
	Type             string        `json:"@type"`
	SourcePort       string        `json:"source_port"`
	SourceChannel    string        `json:"source_channel"`
	Token            Token         `json:"token"`
	Sender           string        `json:"sender"`
	Receiver         string        `json:"receiver"`
	TimeoutHeight    TimeoutHeight `json:"timeout_height"`
	TimeoutTimestamp string        `json:"timeout_timestamp"`
	Memo             string        `json:"memo"`
}

type Token struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type TimeoutHeight struct {
	RevisionNumber string `json:"revision_number"`
	RevisionHeight string `json:"revision_height"`
}

type AuthInfo struct {
	SignerInfos []interface{} `json:"signer_infos"`
	Fee         Fee           `json:"fee"`
}

type Fee struct {
	Amount   []Token `json:"amount"`
	GasLimit string  `json:"gas_limit"`
	Payer    string  `json:"payer"`
	Granter  string  `json:"granter"`
}

type Config struct {
	SuccessfulNodes []string `toml:"successfulNodes"`
}