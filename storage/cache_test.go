package storage

import (
	"context"
	"fmt"
	"testing"

	"github.com/k0yote/dummy-wallet/util"
	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	path := util.CurrentDir()
	config, err := util.LoadConfig(path)
	require.NoError(t, err)

	cache := NewRedisClient(config)
	require.NotNil(t, cache)

	t.Cleanup(func() {
		cache.Client().FlushAll(context.Background())
		CloseRedisConnection(cache)
	})

	cmd := cache.Set(context.Background(), "test", []byte("test"))
	require.NoError(t, cmd.Err())

	val := cache.Get(context.Background(), "test")
	require.NoError(t, val.Err())

	fmt.Printf("val: %s\n", val.Val())
}
