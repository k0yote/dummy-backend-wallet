package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/k0yote/dummy-wallet/domain"
)

type HeartBeatController struct {
	HearbeatUsecase domain.HeartbeatUsecase
}

func (ctr *HeartBeatController) Ping(c *gin.Context) {
	result := ctr.HearbeatUsecase.Ping()
	resp := NewAPIResponse(EMPRTY_MSG, result)
	c.JSON(200, resp)
}
