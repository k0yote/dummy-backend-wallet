package token

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type payload struct {
	ID       xid.ID `json:"sid"`
	Issuer   string `json:"iss"`
	IssuedAt int64  `json:"iat"`
	Subject  string `json:"sub"`
	Audience xid.ID `json:"aud"`
	Expiry   int64  `json:"exp"`
}

/*
	"sid": xid.New().String(),
	"iss": config.TOTPIssuer,
	"iat": iat.Unix(),
	"aud": xid.New().String(),
	"sub": identifier,
	"exp": iat.Add(time.Hour * 24).Unix(),
*/

func NewPayload(issuer string, subject string) *payload {
	return &payload{
		ID:       xid.New(),
		Issuer:   issuer,
		Subject:  subject,
		IssuedAt: time.Now().Unix(),
		Audience: xid.New(),
		Expiry:   time.Now().Add(time.Hour * 24).Unix(),
	}
}

func (p *payload) Valid() error {
	if time.Now().Unix() > p.Expiry {
		return ErrExpiredToken
	}
	return nil
}
