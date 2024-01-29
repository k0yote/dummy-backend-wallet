package util

import (
	"encoding/base64"

	"github.com/google/uuid"
	"github.com/hashicorp/vault/shamir"
	"github.com/k0yote/dummy-wallet/types"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
)

func GenerateKeyInfo(config Config) (types.KeyInfo, error) {

	entropy, err := GenerateRandomBytes(config, 16)
	if err != nil {
		return types.KeyInfo{}, err
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return types.KeyInfo{}, err
	}

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return types.KeyInfo{}, err
	}

	path := hdwallet.MustParseDerivationPath(hdwallet.DefaultBaseDerivationPath.String())

	account, err := wallet.Derive(path, false)
	if err != nil {
		return types.KeyInfo{}, err
	}

	privateKey, err := wallet.PrivateKeyHex(account)
	if err != nil {
		return types.KeyInfo{}, err
	}

	return types.KeyInfo{
		PrivateKey: privateKey,
		PublicKey:  uuid.New().String(),
		AccountID:  account.Address.Hex(),
	}, nil
}

func GenerateSharedSecret(config Config, privateKey string) ([]string, error) {
	byteShares, err := shamir.Split([]byte(privateKey), config.KeySecretShares, config.KeyThresholdShares)
	if err != nil {
		return nil, err
	}

	var strShares []string
	for _, byteShare := range byteShares {
		strShares = append(strShares, base64.StdEncoding.EncodeToString(byteShare))
	}

	return strShares, nil
}

func CombineThresholdShares(strShares []string, privateKey string) ([]byte, error) {
	byteShares := [][]byte{}
	for _, strShare := range strShares {
		if strShare == "" {
			continue
		}

		byteShare, err := base64.StdEncoding.DecodeString(strShare)
		if err != nil {
			return nil, err
		}

		byteShares = append(byteShares, byteShare)
	}

	return shamir.Combine(byteShares)
}
