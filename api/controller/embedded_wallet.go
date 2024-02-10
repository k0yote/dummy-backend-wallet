package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k0yote/dummy-wallet/domain"
	"github.com/k0yote/dummy-wallet/types"
)

type EmbeddedWalletController struct {
	EmbeddedWalletUsecase domain.EmbeddedWalletUsecase
}

func (ctr *EmbeddedWalletController) Initialize(c *gin.Context) {
	var request types.EmbeddedWalletInitializeRequest

	if err := c.ShouldBind(&request); err != nil {
		resp := NewAPIResponse(err.Error(), nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if err := ctr.EmbeddedWalletUsecase.Initialize(request.Identifier); err != nil {
		resp := NewAPIResponse(err.Error(), nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := NewAPIResponse(EMPRTY_MSG, types.InitializeResponse{Success: true})
	c.JSON(http.StatusOK, resp)
}
