package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) heartbeat(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
