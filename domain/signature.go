package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionSignatures = "signatures"
)

type Signautre struct {
	ID        primitive.ObjectID `bson:"_id" json:"-"`
	Nonce     string             `bson:"nonce" json:"nonce"`
	ExpiredAt int64              `bson:"expiredAt" json:"expiredAt"`
}
