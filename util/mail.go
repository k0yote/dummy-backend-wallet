package util

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
)

func SendEmail(config Config, to string, passcode string) error {
	e := email.NewEmail()
	e.From = config.EmailSenderAddress
	e.To = []string{to}
	e.Subject = fmt.Sprintf("%v is your login code for Dummy Auth Demo", passcode)
	e.HTML = []byte(strings.Replace(mailTemplate, "%v", passcode, 1))

	smtpAuth := smtp.PlainAuth("", config.EmailSenderAddress, config.EmailSenderPassword, config.SmtpAuthAddress)
	return e.Send(config.SmtpServerAddress, smtpAuth)
}
