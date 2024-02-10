package usecase

import (
	"context"
	"time"

	"github.com/k0yote/dummy-wallet/container"
	"github.com/k0yote/dummy-wallet/domain"
)

type heartbeatUsecase struct {
	c container.Container
}

func NewHeartbeatUsecase(c container.Container) domain.HeartbeatUsecase {
	return &heartbeatUsecase{
		c: c,
	}
}

func (hu *heartbeatUsecase) Ping() *domain.HeartbeatResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var redisPing, mongoPing bool = true, true

	_, err := hu.c.GetRedis().Client().Ping(ctx).Result()
	if err != nil {
		redisPing = false
	}

	err = hu.c.GetMongoDB().Client().Ping()
	if err != nil {
		mongoPing = false
	}

	return &domain.HeartbeatResponse{
		Redis: redisPing,
		Mongo: mongoPing,
	}
}
