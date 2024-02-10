package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/k0yote/dummy-wallet/util"
	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Client() *redis.Client
	Set(ctx context.Context, key string, v []byte) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	Del(ctx context.Context, key ...string) *redis.IntCmd
}

type cache struct {
	db *redis.Client
}

func NewRedisClient(c util.Config) (Cache, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Printf("redis host: %s\n", c.RedisHost)
	fmt.Printf("redis port: %s\n", c.RedisPort)
	fmt.Printf("redis password: %s\n", c.RedisPassword)
	fmt.Printf("redis dbname: %d\n", c.RedisDbname)
	rdb := redis.NewClient(&redis.Options{
		Network:      "tcp",
		Addr:         c.RedisHost + ":" + c.RedisPort,
		Password:     c.RedisPassword,
		DB:           c.RedisDbname,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 2 * time.Second,
		PoolTimeout:  2 * time.Minute,
		PoolSize:     1000,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return &cache{}, err
	}

	return &cache{
		db: rdb,
	}, nil
}

func (c *cache) Client() *redis.Client {
	return c.db
}

func (c *cache) Set(ctx context.Context, k string, v []byte) *redis.StatusCmd {
	return c.db.Set(ctx, k, v, -1)
}

func (c *cache) Get(ctx context.Context, key string) *redis.StringCmd {
	return c.db.Get(ctx, key)
}

func (c *cache) Del(ctx context.Context, key ...string) *redis.IntCmd {
	return c.db.Del(ctx, key...)
}
