package token

import (
	"os"

	"github.com/rs/zerolog/log"
)

type JWTMaker struct {
	privateKey []byte
	publicKey  []byte
}

func NewJWTMaker(privKeyPath, pubKeyPath string) (Maker, error) {
	pemKey, err := os.ReadFile(privKeyPath)
	if err != nil {
		log.Error().Err(err).Msg("failed to read private key")
		return nil, err
	}

	pubKey, err := os.ReadFile(pubKeyPath)
	if err != nil {
		log.Error().Err(err).Msg("failed to read private key")
		return nil, err
	}
	return &JWTMaker{
		privateKey: pemKey,
		publicKey:  pubKey,
	}, nil
}
