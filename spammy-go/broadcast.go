package main

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"

	"github.com/cosmos/ibc-go/v4/modules/apps/transfer"
	"github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v4/testing/simapp"
)

func sendIBCTransferViaRPC(senderKeyName, rpcEndpoint string, sequence uint64) (response, txbody string, err error) {
	encodingConfig := simapp.MakeTestEncodingConfig()
	//	initClientCtx := client.Context{}.
	//		WithCodec(encodingConfig.Marshaler).
	//		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
	//		WithTxConfig(encodingConfig.TxConfig).
	//		WithLegacyAmino(encodingConfig.Amino).
	//		WithInput(os.Stdin).
	//		WithHomeDir(simapp.DefaultNodeHome).
	//		WithViper("") // In simapp, we don't use any prefix for env variables.

	// Register IBC and other necessary types
	transferModule := transfer.AppModuleBasic{}
	transferModule.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	// Create a new TxBuilder.
	txBuilder := encodingConfig.TxConfig.NewTxBuilder()

	// Create a new keyring to access keys
	kr, err := keyring.New("cosmos", keyring.BackendTest, "/root/.gaia-rs", nil)
	if err != nil {
		return "", "", err
	}

	info, err := kr.Key(senderKeyName)
	if err != nil {
		return "", "", err
	}

	address := info.GetAddress()
	receiver, _ := generateRandomString(30)
	token := sdk.NewCoin("uatom", sdk.NewInt(1))
	msg := types.NewMsgTransfer(
		"transfer",
		"channel-51",
		token,
		address.String(),
		receiver,
		clienttypes.NewHeight(0, 10000),
		types.DefaultRelativePacketTimeoutTimestamp,
	)

	err = txBuilder.SetMsgs(msg)
	if err != nil {
		return "", "", err
	}

	if len(txBuilder.GetTx().GetMsgs()) == 0 {
		return "", "", fmt.Errorf("transaction has no messages set")
	}

	sigBytes, _, err := kr.SignByAddress(address, msg.GetSignBytes())
	if err != nil {
		fmt.Println("coulnd't sign")
		return "", "", err
	}
	fmt.Println("signed")

	sig := signing.SignatureV2{
		PubKey:   info.GetPubKey(),
		Data:     &signing.SingleSignatureData{SignMode: signing.SignMode_SIGN_MODE_DIRECT, Signature: sigBytes},
		Sequence: sequence,
	}

	err = txBuilder.SetSignatures(sig)
	if err != nil {
		fmt.Println("cannot set signatures")
		panic(err)
	}
	fmt.Println("set signatures")

	// Generate a JSON string.
	txJSONBytes, err := encodingConfig.TxConfig.TxJSONEncoder()(txBuilder.GetTx())
	if err != nil {
		fmt.Println("some issue with the string")
		fmt.Println(err)
		return "", "", err
	}

	resp, err := BroadcastTransaction(txJSONBytes, rpcEndpoint)
	if err != nil {
		fmt.Println("some issue with the broadcast, so here's the bytes")
		return "", "", err
	}

	return resp.BroadcastResult.Log, string(txJSONBytes), nil
}

func BroadcastTransaction(txBytes []byte, rpcEndpoint string) (*BroadcastResponse, error) {
	encodedTx := hex.EncodeToString(txBytes)

	broadcastReq := BroadcastRequest{
		Jsonrpc: "2.0",
		ID:      "",
		Method:  "broadcast_tx_sync",
		BroadcastRequestParams: BroadcastRequestParams{
			Tx: encodedTx,
		},
	}

	reqBytes, err := json.Marshal(broadcastReq)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(rpcEndpoint, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var broadcastResp BroadcastResponse
	err = json.Unmarshal(body, &broadcastResp)
	if err != nil {
		return nil, err
	}

	return &broadcastResp, nil
}

func generateRandomString(sizeKB int) (string, error) {
	// Calculate the number of bytes to generate (2 characters per byte in hex encoding)
	nBytes := sizeKB * 1024 / 2

	randomBytes := make([]byte, nBytes)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(randomBytes), nil
}