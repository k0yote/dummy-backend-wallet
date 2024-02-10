package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/k0yote/dummy-wallet/container"
	"github.com/k0yote/dummy-wallet/domain"
	"github.com/k0yote/dummy-wallet/storage"
	"github.com/k0yote/dummy-wallet/util"
	"github.com/stretchr/testify/require"
)

func TestGenerateReadonlyWallet(t *testing.T) {
	path := util.CurrentDir()
	config, err := util.LoadConfig(path)
	require.NoError(t, err)

	db := storage.NewMongoDatabase(config)
	cache := storage.NewRedisClient(config)

	t.Cleanup(func() {
		db.Database(config.MongoDBDbname).Collection(domain.CollectionReadonlyWallet).Drop(context.Background())
		storage.CloseMongoDBConnection(db.Database(config.MongoDBDbname).Client())
		storage.CloseRedisConnection(cache)
	})

	container := container.NewContainer(db.Database(config.MongoDBDbname), cache, config)

	uc := NewReadonlyUsecase(container)
	result, err := uc.GenerateReadonlyWallet(uuid.New().String())
	require.NoError(t, err)

	fmt.Printf("result: %+v\n", result)

}
