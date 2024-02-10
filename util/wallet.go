package util

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/hashicorp/vault/shamir"
	"github.com/k0yote/dummy-wallet/pkg/hdwallet"
	"github.com/k0yote/dummy-wallet/types"
	"github.com/rs/zerolog/log"
	"github.com/tyler-smith/go-bip39"
)

func GenerateEmbeddedWallet(config Config) (types.KeyInfo, error) {

	entropy, err := entropy(config, 16)
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
		Address:    account.Address.Hex(),
		AccountID:  uuid.New().String(),
	}, nil
}

func GenerateSharedSecret(config Config, privateKey []byte) ([]string, error) {
	fmt.Printf("KeySecretShares: %d\t KeyThresholdShares:%d\n", config.KeySecretShares, config.KeyThresholdShares)
	byteShares, err := shamir.Split(privateKey, config.KeySecretShares, config.KeyThresholdShares)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Private Key: %s\n", hex.EncodeToString(privateKey))

	var strShares []string
	for _, byteShare := range byteShares {
		strShares = append(strShares, hex.EncodeToString(byteShare))
	}

	return strShares, nil
}

func CombineThresholdShares(strShares []string) ([]byte, error) {
	byteShares := [][]byte{}
	for _, strShare := range strShares {
		if strShare == "" {
			continue
		}

		byteShare, err := hex.DecodeString(strShare)
		if err != nil {
			return nil, err
		}

		byteShares = append(byteShares, byteShare)
	}

	return shamir.Combine(byteShares)
}

func GenerateOnetimeWallet(uuid string) (types.KeyInfo, error) {
	var keyInfo types.KeyInfo

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot generate private key")
		return keyInfo, err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal().Err(err).Msg("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return keyInfo, err
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return types.KeyInfo{
		PrivateKey: hex.EncodeToString(privateKeyBytes),
		Address:    address,
		AccountID:  uuid,
	}, nil
}

func entropy(config Config, length int) ([]byte, error) {
	ctx := context.Background()
	client, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// Build the request.
	req := &kmspb.GenerateRandomBytesRequest{
		Location:        config.KmsResourceLocation,
		LengthBytes:     int32(length),
		ProtectionLevel: kmspb.ProtectionLevel_HSM,
	}

	result, err := client.GenerateRandomBytes(ctx, req)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}
