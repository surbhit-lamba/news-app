package router

import (
	"context"
	"net/http"
	"news-app/internal/env"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func SetupRouter(ctx context.Context, ev *env.Env, nrApp *newrelic.Application) *gin.Engine {
	r := gin.New()

	r.Use(nrgin.Middleware(nrApp))

	r.Use(env.MiddleWare(ev))

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "welcome to news-app")
	})

	setRoutes(r.Group("stocks"), ctx)

	return r
}

func setRoutes(r *gin.RouterGroup, ctx context.Context) {
	setHealthRoutes(r, ctx)
}

func setHealthRoutes(r *gin.RouterGroup, ctx context.Context) {
	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "yoo healthy!!")
	})
}
