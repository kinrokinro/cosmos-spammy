package main

import (
	"fmt"

	bip39 "github.com/tyler-smith/go-bip39"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types"
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
	bech32PubKey, err := bech32ifyPubKeyWithCustomHRP("celestia", pubKey)
	if err != nil {
		panic(err)
	}

	fmt.Println("Address Ought to be", bech32PubKey)

	return privKey, pubKey, bech32PubKey
}

func bech32ifyPubKeyWithCustomHRP(hrp string, pubKey cryptotypes.PubKey) (string, error) {
	return types.Bech32ifyAddressBytes(hrp, pubKey.Bytes())
}
