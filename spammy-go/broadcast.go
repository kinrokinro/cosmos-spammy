package main

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
)

func sendIBCTransferViaRPC(senderKeyName, rpcEndpoint string, sequence uint64) (broadcastlog string, err error) {
	encodingConfig := simapp.MakeTestEncodingConfig()

	// Create a new TxBuilder.
	txBuilder := encodingConfig.TxConfig.NewTxBuilder()

	// Create a new keyring to access keys
	kr, err := keyring.New("cosmos", keyring.BackendTest, "/root/.gaia-rs", nil)
	if err != nil {
		return "", err
	}

	info, err := kr.Key(senderKeyName)
	if err != nil {
		return "", err
	}

	address := info.GetAddress()
	reciever, err := generateRandomString(30)
	token := sdk.NewCoin("uatom", sdk.NewInt(1))
	msg := types.NewMsgTransfer(
		"transfer",
		"channel-51",
		token,
		address.String(),
		reciever,
		clienttypes.NewHeight(0, 10000),
		types.DefaultRelativePacketTimeoutTimestamp,
	)

	err = txBuilder.SetMsgs(msg)
	if err != nil {
		return "", err
	}

	sigBytes, _, err := kr.SignByAddress(address, msg.GetSignBytes())
	if err != nil {
		return "", err
	}

	sig := signing.SignatureV2{
		PubKey:   info.GetPubKey(),
		Data:     &signing.SingleSignatureData{SignMode: signing.SignMode_SIGN_MODE_DIRECT, Signature: sigBytes},
		Sequence: sequence,
	}

	txBuilder.SetSignatures(sig)

	bz, err := encodingConfig.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return "", err
	}

	broadcastReq := map[string]interface{}{
		"tx": hex.EncodeToString(bz),
	}
	reqBytes, err := json.Marshal(broadcastReq)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(rpcEndpoint+"/broadcast_tx_sync", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
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
