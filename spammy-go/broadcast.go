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
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
)

func sendIBCTransferViaRPC(senderKeyName, rpcEndpoint string, sequence uint64) (broadcastlog string, err error) {
	encodingConfig := simapp.MakeTestEncodingConfig()

	// Create a new TxBuilder.
	txBuilder := encodingConfig.TxConfig.NewTxBuilder()

	kr, err := keyring.New("cosmos", keyring.BackendTest, "/root/.gaia-rs", nil) // adjust paths/backend as necessary
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

	signerData := authsigning.SignerData{
		ChainID:       "provider",
		AccountNumber: 490,      // set actual account number
		Sequence:      sequence, // set actual sequence number
	}

	err = txBuilder.SetMsgs(msg)
	if err != nil {
		return "", err
	}

	// First round: we gather all the signer infos. We use the "set empty
	// signature" hack to do that.
	var sigsV2 []signing.SignatureV2
	sigV2 := signing.SignatureV2{
		PubKey: info.GetPubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  encodingConfig.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: 69,
	}

	sigsV2 = append(sigsV2, sigV2)
	err = txBuilder.SetSignatures(sigsV2...)
	if err != nil {
		return "", err
	}

	// Second round: all signer infos are set, so each signer can sign.
	sigsV2 = []signing.SignatureV2{}
	signed, pubkey, err := kr.Sign("test", msg.GetSignBytes())
	if err != nil {
		return "", err
	}

	fee := sdk.NewCoin("uatom", sdk.NewInt(5000))
	fees := sdk.NewCoins(fee)

	txBuilder.SetGasLimit(300000)
	txBuilder.SetFeeAmount(fees)
	txBuilder.SetMemo("testing 1 2 3")

	err = txBuilder.SetSignatures(sigV2)
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

	fmt.Println("Response:", string(body))
	return broadcastlog, nil
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
