package env

import (
	"context"
	"news-app/pkg/httpclient"

	"github.com/gin-gonic/gin"
)

type Env struct {
	webHTTPRequestClient *httpclient.RequestClient
}

const (
	// EnvCtxKey is the key to set and retrieve Env in context
	EnvCtxKey string = "env"
)

// NewEnv returns a new Env instance
func NewEnv(options ...func(env *Env)) *Env {
	env := &Env{}

	for _, option := range options {
		option(env)
	}

	return env
}

func MiddleWare(env *Env) func(*gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Set(EnvCtxKey, env)
		ctx.Next()
	}
}

// WithContext returns a context containing the env Value
func (env *Env) WithContext(ctx context.Context) context.Context {
	nctx := context.WithValue(ctx, EnvCtxKey, env)
	return nctx
}

func FromContext(ctx context.Context) *Env {
	env, ok := ctx.Value(EnvCtxKey).(*Env)
	if !ok {
		panic("could not fetch env from context")
	}

	return env
}

// WithAlphaVantageHttpConn sets alpha vantage http client in the Env
func WithAlphaVantageHttpConn(webHTTPRequestClient *httpclient.RequestClient) func(*Env) {
	return func(env *Env) {
		env.webHTTPRequestClient = webHTTPRequestClient
	}
}

func (env *Env) AlphaVantageHttpConn() *httpclient.RequestClient {
	return env.webHTTPRequestClient
}
