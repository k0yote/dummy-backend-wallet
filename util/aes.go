package util

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
)

var keyPhrase []byte

func Encrypt(config Config, plain string) (string, error) {
	keyPhrase = []byte(config.KeyStringPhrase)
	buf := new(bytes.Buffer)
	if err := gob.NewEncoder(buf).Encode(plain); err != nil {
		return "", err
	}

	encByte, err := encrypt(buf.Bytes())
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encByte), nil
}

func Decrypt(config Config, enc string) (string, error) {
	keyPhrase = []byte(config.KeyStringPhrase)
	b, err := base64.StdEncoding.DecodeString(enc)
	if err != nil {
		return "", err
	}

	decByte, err := encrypt(b)
	if err != nil {
		return "", err
	}

	var dec string

	if err := gob.NewDecoder(bytes.NewReader(decByte)).Decode(&dec); err != nil {
		return "", err
	}

	return dec, nil
}

func encrypt(payload []byte) ([]byte, error) {
	encOutput := make([]byte, len(payload))
	for i := 0; i < len(payload); i++ {
		encOutput[i] = payload[i] ^ keyPhrase[i%len(keyPhrase)]
	}

	return encOutput, nil
}
