package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type PingController struct{}

func NewPingController() *PingController {
	return &PingController{}
}

func (controller *PingController) Routes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/ping", controller.Ping)
}

func (controller *PingController) Ping(c *gin.Context) {
	log.Info().Msg("Handling Request GET Ping")

	SendOKWithResult(c, "pong")
}
