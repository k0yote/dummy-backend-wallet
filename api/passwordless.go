package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k0yote/dummy-wallet/util"
)

type initRequest struct {
	EmailAddress string `json:"emailAddress" binding:"required"`
}

type initResponse struct {
	Success bool `json:"success"`
}

func (server *Server) initialize(ctx *gin.Context) {
	var req initRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	passcode, err := util.GetPassCode(server.config, req.EmailAddress, 600)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if err := util.SendEmail(server.config, req.EmailAddress, passcode); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := initResponse{
		Success: true,
	}

	ctx.JSON(http.StatusOK, resp)
}

type authenticateRequest struct {
	EmailAddress string `json:"emailAddress" binding:"required"`
}

type authenticateResponse struct {
	Success bool `json:"success"`
}

func (server *Server) authenticate(ctx *gin.Context) {

}
