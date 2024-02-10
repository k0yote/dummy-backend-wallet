package route

import (
	"github.com/gin-gonic/gin"
	"github.com/k0yote/dummy-wallet/container"
)

func NewEmbeddedWalletRouter(c container.Container, group *gin.RouterGroup) {
	// ctr := &controller.EmbeddedWalletController{
	// 	EmbeddedWalletUsecase: usecase.NewEmbeddedWalletUsecase(c),
	// }

}
