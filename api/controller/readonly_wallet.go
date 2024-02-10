package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k0yote/dummy-wallet/domain"
	"github.com/k0yote/dummy-wallet/types"
)

type ReadonlyWalletController struct {
	ReadonlyWalletUsecase domain.ReadonlyWalletUsecase
}

func (ctr *ReadonlyWalletController) Generate(c *gin.Context) {
	var request types.GenerateReadonlyWalletRequest
	if err := c.ShouldBind(&request); err != nil {
		resp := NewAPIResponse(err.Error(), nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	ret, err := ctr.ReadonlyWalletUsecase.GenerateReadonlyWallet(request.AccountID)
	if err != nil {
		resp := NewAPIResponse(err.Error(), nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := NewAPIResponse(EMPRTY_MSG, ret)
	c.JSON(http.StatusOK, resp)
}

func (ctr *ReadonlyWalletController) Search(c *gin.Context) {
	if len(c.Param("type")) == 0 {
		resp := NewAPIResponse("Invalid type", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if len(c.Param("id")) == 0 {
		resp := NewAPIResponse("Invalid id", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	searchType := types.SearchType(c.Param("type"))
	if !searchType.IsValid() {
		resp := NewAPIResponse("Invalid type", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	ret, err := ctr.ReadonlyWalletUsecase.FindBy(searchType, c.Param("id"))
	if err != nil {
		resp := NewAPIResponse(err.Error(), nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := NewAPIResponse(EMPRTY_MSG, ret)
	c.JSON(http.StatusOK, resp)
}
