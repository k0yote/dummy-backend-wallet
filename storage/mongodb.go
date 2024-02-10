package storage

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/k0yote/dummy-wallet/storage/mongodb"
	"github.com/k0yote/dummy-wallet/util"
)

func NewMongoDatabase(conf util.Config) mongodb.Client {
	mongodbURI, err := getConnectionString(conf)
	if err != nil {
		log.Fatal().Err(err).Msg("Error getting Mongo database connection string")
	}

	client, err := mongodb.NewClient(mongodbURI)
	if err != nil {
		log.Fatal().Err(err).Msg("Error getting Mongo database client")
	}

	if err = client.Ping(); err != nil {
		log.Fatal().Err(err).Msg("Error getting Mongo database ping")
	}

	log.Info().Msg("Successfully mongo database connection")

	return client
}

func CloseMongoDBConnection(client mongodb.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal().Err(err).Msg("Error closing MongoDB connection")
	}

	log.Info().Msg("Connection to MongoDB closed.")
}

func getConnectionString(c util.Config) (string, error) {

	dbProtocol := c.MongoDBProtocol
	dbHost := c.MongoDBHost
	dbUser := c.MongoDBUsername
	dbPass := c.MongoDBPassword
	dbName := c.MongoDBDbname
	dbPort := c.MongoDBPort

	switch dbProtocol {
	case "mongodb":
		connStr := fmt.Sprintf("%s://%s:%s@%s:%s", dbProtocol, dbUser, dbPass, dbHost, dbPort)
		if dbUser == "" || dbPass == "" {
			connStr = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
		}
		return connStr, nil
	case "mongodb+srv":
		connStr := fmt.Sprintf("%s://%s:%s@%s/%s", dbProtocol, dbUser, dbPass, dbHost, dbName)
		if dbUser == "" || dbPass == "" {
			connStr = fmt.Sprintf("%s://%s/%s", dbProtocol, dbHost, dbName)
		}
		return connStr, nil
	default:
		return "", fmt.Errorf("protocol not supported: %v", dbProtocol)
	}
}
