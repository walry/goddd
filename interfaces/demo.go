package interfaces

import (
	"github.com/gin-gonic/gin"
	"youras/application/factory"
	"youras/interfaces/controller"
)

func registerDemo(handler *gin.Engine, f *factory.Factory) {
	demo := controller.NewDemoController(f.CreateDemoService())
	handler.GET("demo/:id", demo.Query)
	handler.POST("demo", demo.Update)
}
