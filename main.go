package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/k0yote/dummy-wallet/api/route"
	"github.com/k0yote/dummy-wallet/bootstrap"
	"github.com/k0yote/dummy-wallet/container"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	app := bootstrap.App()
	mongo := app.Mongo.Database(app.Config.MongoDBDbname)
	container := container.NewContainer(mongo, app.Redis, app.Config)
	defer app.CloseDBConnection()

	if app.Config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	handler := gin.Default()

	route.Setup(container, handler)

	srv := &http.Server{
		Addr:    app.Config.HTTPServerAddress,
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msgf("cannot start server on address: %v\n", app.Config.HTTPServerAddress)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msgf("Shutting down server...: %v\n", app.Config.HTTPServerAddress)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msgf("Server forced to shutdown: %v\n", err)
	}
}
