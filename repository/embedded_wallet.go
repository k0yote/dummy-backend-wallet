package repository

import (
	"context"
	"time"

	"github.com/k0yote/dummy-wallet/domain"
	"github.com/k0yote/dummy-wallet/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type embeddedWalletRepository struct {
	db         mongodb.MongoDatabase
	collection string
}

func NewEmbeddedWalletRepository(db mongodb.MongoDatabase, collection string) domain.PasswordlessRepository {
	return &passwordlessRepository{
		db:         db,
		collection: collection,
	}
}

func (r *embeddedWalletRepository) Insert(ctx context.Context, data *domain.Passwordless) error {
	_, err := r.db.Collection(r.collection).InsertOne(ctx, data)
	return err
}

func (r *embeddedWalletRepository) FindBy(ctx context.Context, code, email string) (*domain.Passwordless, error) {
	var result domain.Passwordless

	options := options.Find()
	options.SetSort(bson.D{{Key: "_id", Value: -1}}) // Specify the type of the composite literal explicitly
	options.SetLimit(1)

	filter := bson.D{{Key: "code", Value: code}, {Key: "email", Value: email}, {Key: "expiredAt", Value: bson.D{{Key: "$gt", Value: time.Now().Unix()}}}, {Key: "confirmed", Value: false}}

	cursor, err := r.db.Collection(r.collection).Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		err = cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		return &result, nil
	}

	return nil, err
}
