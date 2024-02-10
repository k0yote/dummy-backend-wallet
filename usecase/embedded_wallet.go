package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/k0yote/dummy-wallet/container"
	"github.com/k0yote/dummy-wallet/domain"
	"github.com/k0yote/dummy-wallet/issuer"
	"github.com/k0yote/dummy-wallet/repository"
	"github.com/k0yote/dummy-wallet/types"
	"github.com/k0yote/dummy-wallet/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type embeddedWalletUsecase struct {
	c container.Container
}

func NewEmbeddedWalletUsecase(c container.Container) domain.EmbeddedWalletUsecase {
	return &embeddedWalletUsecase{
		c: c,
	}
}

func (pu *embeddedWalletUsecase) Initialize(email string) error {
	config := pu.c.GetConfig()
	repo := repository.NewPasswordlessRepository(pu.c.GetMongoDB(), domain.CollectionEmbeddedWallet)

	// generate wallet
	keyInfo, err := util.GenerateEmbeddedWallet(config)

	fmt.Printf("keyInfo: %+v\n", keyInfo)
	// otp passcode
	passCode, err := util.IssuePassCode(config, email)
	if err != nil {
		return err
	}

	// send email
	if err := util.SendEmail(config, email, passCode.Code); err != nil {
		return err
	}

	// save to db
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := repo.Insert(ctx, &domain.Passwordless{
		ID:        primitive.NewObjectID(),
		Email:     email,
		Code:      passCode.Code,
		Secret:    passCode.Secret,
		Confirmed: false,
		ExpiredAt: time.Now().Add(time.Duration(config.PassCodeExpirePeriod) * time.Minute).Unix(),
		CreatedAt: time.Now().Unix(),
	}); err != nil {
		return err
	}

	return nil
}

func (pu *embeddedWalletUsecase) Authenticate(code, email string) (*types.User, error) {
	config := pu.c.GetConfig()
	repo := repository.NewPasswordlessRepository(pu.c.GetMongoDB(), domain.CollectionPasswordless)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data, err := repo.FindBy(ctx, code, email)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errRecordNotFound
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

	token, err := util.GenerateJWTToken(email, entity.Identifier, config)
	if err != nil {
		return nil, err
	}

	return &types.User{
		User: types.UserDetail{
			ID:             entity.Identifier,
			CreatedAt:      entity.State.CreatedAt.Unix(),
			LinkedAccounts: []interface{}{},
		},
		IsNewUser:           true,
		Token:               token.AccessToken,
		RefreshToken:        token.RefreshToken,
		SessionUpdateAction: "set",
	}, nil
}
