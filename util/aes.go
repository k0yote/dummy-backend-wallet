package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"encoding/hex"
	"math/big"

	"github.com/k0yote/dummy-wallet/types"
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

func Aes256Encode(plaintext string) (types.EncryptedInfo, error) {
	var encInfo types.EncryptedInfo

	key, err := GenerateRandomString(32)
	if err != nil {
		return encInfo, err
	}

	iv, err := GenerateRandomString(16)
	if err != nil {
		return encInfo, err
	}

	bPlaintext := PKCS5Padding([]byte(plaintext), aes.BlockSize, len(plaintext))
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return encInfo, err
	}

	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, bPlaintext)

	return types.EncryptedInfo{
		CipherText: base64.StdEncoding.EncodeToString(ciphertext),
		EncKey:     key,
		IV:         iv,
	}, nil
}

func Aes256Decode(cipherText string, encKey string, iv string) (decryptedString string) {
	bKey := []byte(encKey)
	bIV := []byte(iv)
	cipherTextDecoded, err := hex.DecodeString(cipherText)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(bKey)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, bIV)
	mode.CryptBlocks([]byte(cipherTextDecoded), []byte(cipherTextDecoded))
	return string(cipherTextDecoded)
}

func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func GenerateRandomStringURLSafe(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), err
}

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-/=+"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}
