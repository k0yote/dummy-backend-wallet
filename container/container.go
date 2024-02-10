package container

import (
	"github.com/k0yote/dummy-wallet/storage/cache"
	"github.com/k0yote/dummy-wallet/storage/mongodb"
	"github.com/k0yote/dummy-wallet/util"
)

type Container interface {
	GetMongoDB() mongodb.MongoDatabase
	GetRedis() cache.Cache
	GetConfig() util.Config
}

type container struct {
	mongo  mongodb.MongoDatabase
	caceh  cache.Cache
	config util.Config
}

func NewContainer(mongo mongodb.MongoDatabase, cache cache.Cache, config util.Config) Container {
	return &container{
		mongo:  mongo,
		caceh:  cache,
		config: config,
	}
}

func (c *container) GetMongoDB() mongodb.MongoDatabase {
	return c.mongo
}

func (c *container) GetRedis() cache.Cache {
	return c.caceh
}

func (c *container) GetConfig() util.Config {
	return c.config
}
