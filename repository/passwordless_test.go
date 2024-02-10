package repository

import (
	"context"
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/k0yote/dummy-wallet/domain"
	"github.com/k0yote/dummy-wallet/storage"
	"github.com/k0yote/dummy-wallet/util"
	"github.com/rs/xid"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsert(t *testing.T) {
	path := util.CurrentDir()

	config, err := util.LoadConfig(path)
	require.NoError(t, err)

	db := storage.NewMongoDatabase(config)

	repo := NewPasswordlessRepository(db.Database(config.MongoDBDbname), domain.CollectionPasswordless)

	t.Cleanup(func() {
		db.Database(config.MongoDBDbname).Collection(domain.CollectionPasswordless).Drop(context.Background())
		storage.CloseMongoDBConnection(db.Database(config.MongoDBDbname).Client())
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = repo.Insert(ctx, &domain.Passwordless{
		ID:        primitive.NewObjectID(),
		Email:     "aaa@aaa.com",
		Code:      "123456",
		Secret:    "aaaa",
		Confirmed: false,
		ExpiredAt: time.Now().Add(5 * time.Minute).Unix(),
		CreatedAt: time.Now().Unix(),
	})

	require.NoError(t, err)

	data, err := repo.FindBy(ctx, "123457", "aaa@aaa.com")
	require.NoError(t, err)

	fmt.Printf("data: %+v\n", data)

}

func TestDummyWallet(t *testing.T) {
	xuuid := xid.New()
	fmt.Printf("aid: %d\n", len(xuuid))

	for i := 0; i < 12; i++ {
		fmt.Printf("aid: %d\n", xuuid[i])
		fmt.Printf("aid: %s\n", hex.EncodeToString(xid.New().Bytes()))
	}

	generatedId := uuid.New()
	fmt.Printf("aid: %d\n", len(generatedId))

	id := [16]byte{}
	for i := 0; i < 16; i++ {
		id[i] = generatedId[i]
	}

	fmt.Printf("aid: %s\n", hex.EncodeToString(id[:]))
}
