package usecase

import (
	"context"
	"encoding/hex"
	"strings"
	"time"

	"github.com/k0yote/dummy-wallet/container"
	"github.com/k0yote/dummy-wallet/domain"
	"github.com/k0yote/dummy-wallet/issuer"
	"github.com/k0yote/dummy-wallet/repository"
	"github.com/k0yote/dummy-wallet/types"
	"github.com/k0yote/dummy-wallet/util"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type readonlyUsecase struct {
	c container.Container
}

func NewReadonlyUsecase(c container.Container) domain.ReadonlyWalletUsecase {
	return &readonlyUsecase{
		c: c,
	}
}

func (ru *readonlyUsecase) GenerateReadonlyWallet(accountID string) (*types.ReadonlyWalletResponse, error) {
	config := ru.c.GetConfig()
	repo := repository.NewReadonlyWalletRepository(ru.c.GetMongoDB(), domain.CollectionReadonlyWallet)

	keyInfo, err := util.GenerateOnetimeWallet(accountID)
	if err != nil {
		return nil, err
	}

	b, err := hex.DecodeString(keyInfo.PrivateKey)
	if err != nil {
		return nil, err
	}

	shares, err := util.GenerateSharedSecret(ru.c.GetConfig(), b)
	if err != nil {
		return nil, err
	}

	plainText := strings.Join(shares[1:], ":")

	encrypted, err := util.Aes256Encode(plainText)
	if err != nil {
		return nil, err
	}

	issuerNode := issuer.Init(&issuer.RequestOpts{
		BaseURL:  config.IssuerBaseURL,
		Username: config.BasicAuthUsername,
		Password: config.BasicAuthPassword,
	})

	entity, err := issuerNode.CreateEntity(&issuer.CreateEntityArgs{
		DidMetadata: &issuer.DidMetadata{
			Method:     "polygonid",
			Blockchain: "polygon",
			Network:    "mumbai",
			Type:       "BJJ",
		},
	})
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := repo.Insert(ctx, &domain.ReadonlyWallet{
		ID:            primitive.NewObjectID(),
		Identifier:    entity.Identifier,
		AccountID:     accountID,
		Address:       keyInfo.Address,
		RecoveryShare: shares[2],
		BackupShare:   shares[0],
		CreatedAt:     time.Now().Unix(),
	}); err != nil {
		return nil, err
	}

	return &types.ReadonlyWalletResponse{
		Identifier: entity.Identifier,
		Address:    keyInfo.Address,
		AccountID:  accountID,
		CipherText: encrypted.CipherText,
		IV:         encrypted.IV,
		EncKey:     encrypted.EncKey,
		Nonce:      xid.New().String(),
	}, nil
}

func (ru *readonlyUsecase) FindBy(typ types.SearchType, id string) (*domain.ReadonlyWallet, error) {
	repo := repository.NewReadonlyWalletRepository(ru.c.GetMongoDB(), domain.CollectionReadonlyWallet)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return repo.FindBy(ctx, typ, id)
}
