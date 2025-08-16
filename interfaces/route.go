package interfaces

import (
	"github.com/gin-gonic/gin"
	"youras/application/factory"
	"youras/infra/config"
)

func NewHttpHandler(f *factory.Factory, cfg config.AppConfig) *gin.Engine {
	gin.SetMode(cfg.GinMode)
	g := gin.New()
	for _, r := range registers {
		r(g, f)
	}
	return g
}

type register func(g *gin.Engine, f *factory.Factory)

var registers = []register{
	registerPing,
	registerDemo,
}
