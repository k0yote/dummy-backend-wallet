package util

import (
	"crypto/rand"
	"errors"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

var ErrInvalidPassCode = errors.New("passcode is invalid")

func GetPassCode(config Config, email string, period uint) (string, error) {

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      config.TOTPIssuer,
		AccountName: email,
		Period:      period,
		Digits:      otp.DigitsSix,
		SecretSize:  20,
		Secret:      []byte{},
		Algorithm:   otp.AlgorithmSHA512,
		Rand:        rand.Reader,
	})
	if err != nil {
		return "", err
	}

	passcode, err := generatePassCode(period, key.Secret())
	if err != nil {
		return "", err
	}

	valid, err := validatePassCode(period, key.Secret(), passcode)
	if err != nil {
		return "", err
	}

	if !valid {
		return "", ErrInvalidPassCode
	}

	return passcode, nil
}

func generatePassCode(period uint, secret string) (string, error) {
	return totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period:    period,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA512,
	})
}

func validatePassCode(period uint, secret string, passcode string) (bool, error) {
	return totp.ValidateCustom(passcode, secret, time.Now().UTC(), totp.ValidateOpts{
		Period:    period,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA512,
	})
}
