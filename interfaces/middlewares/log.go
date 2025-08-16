package middlewares

import (
	"github.com/gin-gonic/gin"
	"time"
	"youras/infra/yctx"
	"youras/pkg/ylog"
)

type _logger struct{}

var accessLoggerTemplate = ylog.Access().With(
	"request_id", "",
	"method", "",
	"path", "",
	"status", 0,
	"cost", 0,
)

var globalLoggerTemplate = ylog.Log().With(
	"request_id", "",
	"cost", 0,
	"caller", "")

func AccessLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		myctx := yctx.FromContext(c.Request.Context())
		c.Request.WithContext(yctx.WithContext(c.Request.Context(), myctx))
		c.Next()
		elapsed := time.Since(start)
		accessLoggerTemplate.Infow("",
			"request_id", myctx.TraceId(),
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"cost", elapsed.Seconds(),
		)
	}
}
