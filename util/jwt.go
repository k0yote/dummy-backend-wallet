package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/k0yote/dummy-wallet/types"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
)

var kid = "dummy-wallet-key"

func GenerateJWTToken(email, identifier string, config Config) (types.JWTToken, error) {

	var token types.JWTToken

	pemKey, err := os.ReadFile(config.PrivkeyPath)
	if err != nil {
		log.Error().Err(err).Msg("failed to read private key")
		return token, err
	}

	var ecdsaKey *ecdsa.PrivateKey

	if ecdsaKey, err = jwt.ParseECPrivateKeyFromPEM([]byte(pemKey)); err != nil {
		log.Error().Err(err).Msg("failed to parse private key")
		return token, err
	}

	iat := time.Now()

	claims := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sid": xid.New().String(),
		"iss": config.TOTPIssuer,
		"iat": iat.Unix(),
		"aud": xid.New().String(),
		"sub": identifier,
		"exp": iat.Add(time.Hour * 24).Unix(),
	})

	claims.Header["kid"] = kid

	accessToken, err := claims.SignedString(ecdsaKey)
	if err != nil {
		log.Error().Err(err).Msg("failed to sign token")
		return token, err
	}

	refreshToken, err := createRefreshToken(email)
	if err != nil {
		log.Error().Err(err).Msg("failed to create refresh token")
		return token, err
	}

	return types.JWTToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func GenerateRefeshToken(config Config) {

}

func VerifyAccessToken(config Config, accessToken string) {
	pemKey, err := os.ReadFile(config.PubKeyPath)
	if err != nil {
		panic(err)
	}

	key, err := jwt.ParseECPublicKeyFromPEM([]byte(pemKey))
	if err != nil {
		panic(err)
	}

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", token)
}

func ValidateRefreshToken(refreshToken string) error {
	sha1 := sha1.New()
	io.WriteString(sha1, os.Getenv("SECRET_KEY"))

	salt := string(sha1.Sum(nil))[0:16]
	block, err := aes.NewCipher([]byte(salt))
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	data, err := base64.URLEncoding.DecodeString(refreshToken)
	if err != nil {
		return err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plain, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	if string(plain) != refreshToken {
		return errors.New("invalid token")
	}

	claims := jwt.MapClaims{}
	parser := jwt.Parser{}
	token, _, err := parser.ParseUnverified(refreshToken, claims)

	if err != nil {
		return err
	}

	payload, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid token")
	}

	fmt.Printf("%v\n", payload)

	return nil
}

func createRefreshToken(email string) (string, error) {

	sha1 := sha256.New()
	io.WriteString(sha1, os.Getenv("SECRET_KEY"))

	salt := string(sha1.Sum(nil))[0:16]

	block, err := aes.NewCipher([]byte(salt))
	if err != nil {
		log.Error().Err(err).Msg("failed to create new cipher")
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Error().Err(err).Msg("failed to create new GCM")
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		log.Error().Err(err).Msg("failed to read full")
		return "", err
	}

	return base64.URLEncoding.EncodeToString(gcm.Seal(nonce, nonce, []byte(email), nil)), nil
}
