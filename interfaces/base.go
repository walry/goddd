package interfaces

import (
	"github.com/gin-gonic/gin"
	"youras/application/factory"
)

func registerPing(g *gin.Engine, _ *factory.Factory) {
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})
}
