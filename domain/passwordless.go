package domain

import (
	"context"

	"github.com/k0yote/dummy-wallet/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionPasswordless = "passwordless"
)

type Passwordless struct {
	ID        primitive.ObjectID `bson:"_id" json:"-"`
	Email     string             `bson:"email" json:"email"`
	Code      string             `bson:"code" json:"code"`
	Secret    string             `bson:"secret" json:"secret"`
	Confirmed bool               `bson:"confirmed" json:"confirmed"`
	ExpiredAt int64              `bson:"expiredAt" json:"expiredAt"`
	CreatedAt int64              `bson:"createdAt" json:"createdAt"`
}

type PasswordlessUsecase interface {
	Initialize(email string) error
	Authenticate(email, code string) (*types.User, error)
}

type PasswordlessRepository interface {
	Insert(ctx context.Context, data *Passwordless) error
	FindBy(ctx context.Context, code, email string) (*Passwordless, error)
	Update(ctx context.Context, filter, update primitive.D) error
}
