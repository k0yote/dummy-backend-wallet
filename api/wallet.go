package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createWalletRequest struct {
	UUID string `json:"uuid"`
}

type walletResponse struct {
	PrivateKey string `json:"privateKey"`
	AccountID  string `json:"accountId"`
}

func newWalletResponse(privateKey, accountID string) walletResponse {
	return walletResponse{
		PrivateKey: privateKey,
		AccountID:  accountID,
	}
}

func (server *Server) createWallet(ctx *gin.Context) {
	var req createWalletRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//TODO: implement

}
