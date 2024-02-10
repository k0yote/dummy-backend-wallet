package bootstrap

import (
	"github.com/k0yote/dummy-wallet/storage"
	"github.com/k0yote/dummy-wallet/storage/cache"
	"github.com/k0yote/dummy-wallet/storage/mongodb"
	"github.com/k0yote/dummy-wallet/util"
	"github.com/rs/zerolog/log"
)

type Application struct {
	Config util.Config
	Mongo  mongodb.Client
	Redis  cache.Cache
}

func App() Application {
	app := &Application{}

	path := util.CurrentDir()
	conf, err := util.LoadConfig(path)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	app.Mongo = storage.NewMongoDatabase(conf)
	app.Redis = storage.NewRedisClient(conf)
	app.Config = conf

	return *app
}

func (app *Application) CloseDBConnection() {
	storage.CloseMongoDBConnection(app.Mongo)
	storage.CloseRedisConnection(app.Redis)
}
