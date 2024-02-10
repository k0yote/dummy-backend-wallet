package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionLinkedAccounts = "linkedAccounts"
)

type LinkedAccounts struct {
	ID              primitive.ObjectID `bson:"_id" json:"-"`
	InitType        string             `bson:"initType" json:"initType"`
	Email           string             `bson:"email" json:"email"`
	EmailVerifiedAt int64              `bson:"emailVerifiedAt" json:"emailVerifiedAt"`
	AuthType        string             `bson:"authType" json:"authType"`
	Address         string             `bson:"address" json:"address"`
	WalletIndex     int                `bson:"walletIndex" json:"walletIndex"`
	ChainID         string             `bson:"chainId" json:"chainId"`
	ChainType       string             `bson:"chainType" json:"chainType"`
	WalletClient    string             `bson:"walletClient" json:"walletClient"`
	ConnectorType   string             `bson:"connectorType" json:"connectorType"`
	AuthVerifiedAt  int64              `bson:"authVerifiedAt" json:"authVerifiedAt"`
	RecoveryMethod  string             `bson:"recoveryMethod" json:"recoveryMethod"`
}

type LinkedAccountsRepository interface {
	Insert(ctx context.Context, data *Passwordless) error
	FindBy(ctx context.Context, code, email string) (*Passwordless, error)
}
