package route

import (
	"github.com/gin-gonic/gin"
	"github.com/k0yote/dummy-wallet/api/controller"
	"github.com/k0yote/dummy-wallet/container"
	"github.com/k0yote/dummy-wallet/usecase"
)

func NewReadonlyWalletRouter(c container.Container, group *gin.RouterGroup) {
	ctr := &controller.ReadonlyWalletController{
		ReadonlyWalletUsecase: usecase.NewReadonlyUsecase(c),
	}

	group.POST(APIReadonlyWalletGenerate, ctr.Generate)
	group.GET(APIReadonlyWalletSearch, ctr.Search)
}
