package domain

import (
	"context"

	"github.com/k0yote/dummy-wallet/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionEmbeddedWallet = "embeddedWallet"
)

type EmbeddedWallet struct {
	ID            primitive.ObjectID `bson:"_id" json:"-"`
	Identifier    string             `bson:"identifier" json:"identifier"`
	SessionID     string             `bson:"sessionId" json:"sessionId"`
	SecretShare   string             `bson:"secretShare" json:"secretShare"`
	RecoveryShare string             `bson:"recoveryShare" json:"recoveryShare"`
	RecoveryKey   string             `bson:"recoveryKey" json:"recoveryKey"`
	CreatedAt     int64              `bson:"createdAt" json:"createdAt"`
	UpdatedAt     int64              `bson:"updatedAt" json:"updatedAt"`
}

type EmbeddedWalletUsecase interface {
	Initialize(identifier string) error
	Authenticate(email, code string) (*types.User, error)
}

type EmbeddedWalletRepository interface {
	Insert(ctx context.Context, data *Passwordless) error
	FindBy(ctx context.Context, code, email string) (*Passwordless, error)
}
