package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k0yote/dummy-wallet/util"
)

type createWalletRequest struct {
	UUID string `json:"uuid" binding:"required"`
}

type walletResponse struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
	AccountID  string `json:"accountId"`
}

func newWalletResponse(privateKey, publicKey, accountID string) walletResponse {
	return walletResponse{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		AccountID:  accountID,
	}
}

func (server *Server) createWallet(ctx *gin.Context) {
	var req createWalletRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	keyInfo, err := util.GenerateKeyInfo(server.config)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return

	}

	resp := newWalletResponse(keyInfo.PrivateKey, keyInfo.PublicKey, keyInfo.AccountID)
	ctx.JSON(http.StatusOK, resp)
}
