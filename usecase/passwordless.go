package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/k0yote/dummy-wallet/container"
	"github.com/k0yote/dummy-wallet/domain"
	"github.com/k0yote/dummy-wallet/issuer"
	"github.com/k0yote/dummy-wallet/repository"
	"github.com/k0yote/dummy-wallet/types"
	"github.com/k0yote/dummy-wallet/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	errRecordNotFound   = errors.New("record not found")
	errAlreadyConfirmed = errors.New("code already confirmed")
)

type passwordlessUsecase struct {
	c container.Container
}

func NewPasswordlessUsecase(c container.Container) domain.PasswordlessUsecase {
	return &passwordlessUsecase{
		c: c,
	}
}

func (pu *passwordlessUsecase) Initialize(email string) error {
	config := pu.c.GetConfig()
	repo := repository.NewPasswordlessRepository(pu.c.GetMongoDB(), domain.CollectionPasswordless)

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

func (pu *passwordlessUsecase) Authenticate(code, email string) (*types.User, error) {
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

	if data.Confirmed {
		return nil, errAlreadyConfirmed
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

	filter := bson.D{{Key: "_id", Value: data.ID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "confirmed", Value: true}}}}
	if err := repo.Update(ctx, filter, update); err != nil {
		return nil, err
	}

	repo = repository.NewPasswordlessRepository(pu.c.GetMongoDB(), domain.CollectionPasswordless)

	return &types.User{
		User: types.UserDetail{
			ID:        entity.Identifier,
			CreatedAt: entity.State.CreatedAt.Unix(),
			LinkedAccounts: []interface{}{
				types.LinkedEmailAccount{
					Type:       "email",
					Address:    email,
					VerifiedAt: time.Now().Unix(),
				},
			},
		},
		IsNewUser:           true,
		Token:               token.AccessToken,
		RefreshToken:        token.RefreshToken,
		SessionUpdateAction: "set",
	}, nil
}
