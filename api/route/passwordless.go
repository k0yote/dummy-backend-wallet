package route

import (
	"github.com/gin-gonic/gin"
	"github.com/k0yote/dummy-wallet/api/controller"
	"github.com/k0yote/dummy-wallet/container"
	"github.com/k0yote/dummy-wallet/usecase"
)

func NewPasswordlessRouter(c container.Container, group *gin.RouterGroup) {
	ctr := &controller.PasswordlessController{
		PasswordlessUsecase: usecase.NewPasswordlessUsecase(c),
	}

	group.POST(APIEmbeddedWalletInit, ctr.Initialize)
	group.POST(APIEmbeddedWalletSignature, ctr.Authenticate)
}
