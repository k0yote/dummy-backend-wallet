package repository

import (
	"context"

	"github.com/k0yote/dummy-wallet/domain"
	"github.com/k0yote/dummy-wallet/storage/mongodb"
	"github.com/k0yote/dummy-wallet/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type readonlyWalletRepository struct {
	db         mongodb.MongoDatabase
	collection string
}

func NewReadonlyWalletRepository(db mongodb.MongoDatabase, collection string) domain.ReadonlyWalletRepository {
	return &readonlyWalletRepository{
		db:         db,
		collection: collection,
	}
}

func (r *readonlyWalletRepository) Insert(ctx context.Context, data *domain.ReadonlyWallet) error {
	_, err := r.db.Collection(r.collection).InsertOne(ctx, data)
	return err
}

func (r *readonlyWalletRepository) FindBy(ctx context.Context, typ types.SearchType, id string) (*domain.ReadonlyWallet, error) {
	var result domain.ReadonlyWallet

	filter := generateFilter(typ, id)
	if err := r.db.Collection(r.collection).FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func generateFilter(typ types.SearchType, id string) primitive.D {
	key := types.AccountID
	value := id
	if typ == types.Address {
		key = types.Address
	}
	return bson.D{{Key: string(key), Value: value}}
}
