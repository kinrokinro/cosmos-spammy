package main

type BlockResult struct {
	Result struct {
		Block struct {
			Header struct {
				Height string `json:"height"`
			} `json:"header"`
			Data struct {
				Txs []string `json:"txs"`
			} `json:"data"`
		} `json:"block"`
	} `json:"result"`
}

type MempoolResult struct {
	Result struct {
		NTxs string `json:"n_txs"`
	} `json:"result"`
}

type AccountResult struct {
	Account struct {
		Sequence string `json:"sequence"`
	} `json:"account"`
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
