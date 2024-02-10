package util

import (
	"crypto/rand"
	"errors"
	"fmt"
	"net/smtp"
	"strings"
	"time"

	"github.com/jordan-wright/email"
	"github.com/k0yote/dummy-wallet/types"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

var errInvalidPassCode = errors.New("passcode is invalid")

func IssuePassCode(config Config, email string) (types.PasscodeResult, error) {

	var result types.PasscodeResult

	period := uint(config.PassCodeExpirePeriod)
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
		return result, err
	}

	passcode, err := generatePassCode(period, key.Secret())
	if err != nil {
		return result, err
	}

	valid, err := validatePassCode(period, key.Secret(), passcode)
	if err != nil {
		return result, err
	}

	if !valid {
		return result, errInvalidPassCode
	}

	return types.PasscodeResult{
		Code:   passcode,
		Secret: key.Secret(),
	}, nil
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

func SendEmail(config Config, to string, passcode string) error {
	e := email.NewEmail()
	e.From = config.EmailSenderAddress
	e.To = []string{to}
	e.Subject = fmt.Sprintf("%v is your login code for Dummy Auth Demo", passcode)
	e.HTML = []byte(strings.Replace(mailTemplate, "%v", passcode, 1))

	smtpAuth := smtp.PlainAuth("", config.EmailSenderAddress, config.EmailSenderPassword, config.SmtpAuthAddress)
	return e.Send(config.SmtpServerAddress, smtpAuth)
}
