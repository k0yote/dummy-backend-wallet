package route

import (
	"github.com/gin-gonic/gin"
	"github.com/k0yote/dummy-wallet/api/controller"
	"github.com/k0yote/dummy-wallet/container"
	"github.com/k0yote/dummy-wallet/usecase"
)

func NewHeartbeatRouter(c container.Container, group *gin.RouterGroup) {
	ctr := &controller.HeartBeatController{
		HearbeatUsecase: usecase.NewHeartbeatUsecase(c),
	}

	group.GET(APIHeartbeat, ctr.Ping)
}
