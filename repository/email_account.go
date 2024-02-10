package repository

import (
	"github.com/k0yote/dummy-wallet/domain"
	"github.com/k0yote/dummy-wallet/storage/mongodb"
)

type emailAccountRepository struct {
	db         mongodb.MongoDatabase
	collection string
}

func NewEmailAccountRepository(db mongodb.MongoDatabase, collection string) domain.EmailAccountRepository {
	return &emailAccountRepository{
		db:         db,
		collection: collection,
	}
}

func (r *emailAccountRepository) Insert(data *domain.EmailAccount) error {
	return nil
}
