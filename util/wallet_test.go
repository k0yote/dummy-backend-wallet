package util

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestShamirSharingSecret(t *testing.T) {

	path := CurrentDir()
	config, err := LoadConfig(path)
	require.NoError(t, err)

	keyInfo, err := GenerateOnetimeWallet(uuid.New().String())
	require.NoError(t, err)

	b, err := hex.DecodeString(keyInfo.PrivateKey)
	require.NoError(t, err)

	shares, err := GenerateSharedSecret(config, b)
	require.NoError(t, err)

	fmt.Printf("keyInfo: %+v\n", keyInfo)
	fmt.Println("Shares: ", shares)
	fmt.Println("Shares: ", shares[:1])
	fmt.Println("Shares: ", shares[1:])

	plainText := strings.Join(shares[1:], ":")

	p, err := CombineThresholdShares(shares[1:])
	require.NoError(t, err)

	fmt.Println("Private Key: ", hex.EncodeToString(p))

	encrypted, err := Aes256Encode(plainText)
	require.NoError(t, err)
	fmt.Printf("encrypted: %+v\n", encrypted)
}
