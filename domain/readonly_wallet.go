package domain

import (
	"context"

	"github.com/k0yote/dummy-wallet/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionReadonlyWallet = "readonlyWallets"
)

type ReadonlyWallet struct {
	ID            primitive.ObjectID `bson:"_id" json:"-"`
	Identifier    string             `bson:"identifier" json:"identifier"`
	AccountID     string             `bson:"accountId" json:"accountId"`
	Address       string             `bson:"address" json:"address"`
	RecoveryShare string             `bson:"recoveryShare" json:"recoveryShare"`
	BackupShare   string             `bson:"backupShare" json:"backupShare"`
	CreatedAt     int64              `bson:"createdAt" json:"createdAt"`
}

type ReadonlyWalletUsecase interface {
	GenerateReadonlyWallet(uuid string) (*types.ReadonlyWalletResponse, error)
	FindBy(typ types.SearchType, id string) (*ReadonlyWallet, error)
}

type ReadonlyWalletRepository interface {
	Insert(ctx context.Context, data *ReadonlyWallet) error
	FindBy(ctx context.Context, typ types.SearchType, id string) (*ReadonlyWallet, error)
}
