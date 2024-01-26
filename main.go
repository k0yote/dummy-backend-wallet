package main

import (
	"fmt"
	"os"
	"time"

	"github.com/k0yote/dummy-wallet/api"
	"github.com/k0yote/dummy-wallet/util"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func GeneratePassCode(secret string) string {
	passcode, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period:    300,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA512,
	})
	if err != nil {
		panic(err)
	}
	return passcode
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")

	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// FIXME: start from here, move into api
	passcode, err := util.GetPassCode(config, "yjmk0yote@gmail.com", 600)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get passcode")
	}

	util.SendEmail(config, "yjmk0yote@gmail.com", passcode)

	entropy, err := util.GenerateRandomBytes(config, 16)
	if err != nil {
		panic(err)
	}

	keyInfo, err := util.GenerateKeyInfo(entropy)
	if err != nil {
		panic(err)
	}

	fmt.Printf("keyInfo: %+v\n", keyInfo)
	// til to end
	runGinServer(config)
}

func runGinServer(config util.Config) {
	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}
