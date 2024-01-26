package util

import (
	"github.com/google/uuid"
	"github.com/k0yote/dummy-wallet/types"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
)

func GenerateKeyInfo(entropy []byte) (types.KeyInfo, error) {
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
