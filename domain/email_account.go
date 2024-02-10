package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionEmailAccounts = "emailAccounts"
)

type EmailAccount struct {
	ID         primitive.ObjectID `bson:"_id" json:"-"`
	Identifier string             `bson:"identifier" json:"identifier"`
	Email      string             `bson:"email" json:"email"`
	IsVerified bool               `bson:"isVerified" json:"isVerified"`
	VerifiedAt int64              `bson:"verifiedAt" json:"verifiedAt"`
}

type EmailAccountRepository interface {
	Insert(data *EmailAccount) error
}
