package route

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/k0yote/dummy-wallet/container"
)

func Setup(container container.Container, gin *gin.Engine) {
	gin.HandleMethodNotAllowed = true
	setCORSConfig(gin)
	setResponseHeader(gin)
	groupRouter := gin.Group("")
	// All Public APIs
	NewHeartbeatRouter(container, groupRouter)
	NewReadonlyWalletRouter(container, groupRouter)
	NewPasswordlessRouter(container, groupRouter)
}

func setCORSConfig(gin *gin.Engine) {
	gin.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowCredentials: true,
		AllowHeaders:     []string{"*"},
	}))
}

func setResponseHeader(c *gin.Engine) {
	c.Use(func() gin.HandlerFunc {
		return func(c *gin.Context) {
			c.Writer.Header().Set("Content-Type", "application/json")
		}
	}())
}
