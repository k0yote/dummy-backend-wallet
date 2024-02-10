package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k0yote/dummy-wallet/domain"
	"github.com/k0yote/dummy-wallet/types"
)

type PasswordlessController struct {
	PasswordlessUsecase domain.PasswordlessUsecase
}

func (ctr *PasswordlessController) Initialize(c *gin.Context) {
	var request types.InitializeRequest

	if err := c.ShouldBind(&request); err != nil {
		resp := NewAPIResponse(err.Error(), nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if err := ctr.PasswordlessUsecase.Initialize(request.Email); err != nil {
		resp := NewAPIResponse(err.Error(), nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := NewAPIResponse(EMPRTY_MSG, types.InitializeResponse{Success: true})
	c.JSON(http.StatusOK, resp)
}

func (ctr *PasswordlessController) Authenticate(c *gin.Context) {
	var request types.AuthenticateRequest

	if err := c.ShouldBind(&request); err != nil {
		resp := NewAPIResponse(err.Error(), nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	user, err := ctr.PasswordlessUsecase.Authenticate(request.Email, request.Code)
	if err != nil {
		resp := NewAPIResponse(err.Error(), nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := NewAPIResponse(EMPRTY_MSG, user)
	c.JSON(http.StatusOK, resp)
}
