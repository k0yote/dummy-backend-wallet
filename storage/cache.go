package storage

import (
	"github.com/k0yote/dummy-wallet/storage/cache"
	"github.com/k0yote/dummy-wallet/util"
	"github.com/rs/zerolog/log"
)

func NewRedisClient(c util.Config) cache.Cache {
	rdb, err := cache.NewRedisClient(c)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create redis client")
	}

	log.Info().Msg("redis client created")
	return rdb
}

func CloseRedisConnection(cache cache.Cache) {
	if cache.Client() == nil {
		log.Fatal().Msg("redis client is nil")
		return
	}

	if err := cache.Client().Close(); err != nil {
		log.Fatal().Err(err).Msg("failed to close redis connection")
	}

	log.Info().Msg("redis connection closed")
}
