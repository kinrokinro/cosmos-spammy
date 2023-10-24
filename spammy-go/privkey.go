package main

import (
	"fmt"

	bip39 "github.com/tyler-smith/go-bip39"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func getPrivKey(mnemonic []byte) (cryptotypes.PrivKey, cryptotypes.PubKey, string) {
	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(string(mnemonic), "")

	// Create master private key from seed
	masterPriv, chainCode := hd.ComputeMastersFromSeed(seed)

	// Path for Celestia
	// This is just an example path. Adjust based on the actual requirements.
	path := hd.CreateHDPath(118, 0, 0).String()

	derivedPriv, _ := hd.DerivePrivateKeyForPath(masterPriv, chainCode, path)

	privKey := secp256k1.GenPrivKeyFromSecret(derivedPriv)

	pubKey := privKey.PubKey()

	// Convert the public key to Bech32 with custom HRP
	// bech32PubKey, err := bech32ifyPubKeyWithCustomHRP("celestia", pubKey)
	// if err != nil {
	//	panic(err)
	// }

	addressbytes := sdk.AccAddress(pubKey.Address().Bytes())
	address, err := sdk.Bech32ifyAddressBytes("cosmos", addressbytes)
	if err != nil {
		panic(err)
	}

	fmt.Println("Address Ought to be", address)

	return privKey, pubKey, address
}
