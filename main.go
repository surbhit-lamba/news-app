package main

import (
	"context"
	"net/http"
	"news-app/internal/env"
	"news-app/internal/router"
	"news-app/pkg/httpclient"
	"news-app/pkg/log"
	"news-app/pkg/newrelic"
	"os"
	"time"
)

func main() {
	ctx := context.Background()

	nrApp := newrelic.Initialize(
		&newrelic.Options{
			Name:                   os.Getenv("NEWRELIC_NAME"),
			License:                os.Getenv("NEWRELIC_KEY"),
			Enabled:                true,
			CrossApplicationTracer: true,
			DistributedTracer:      true,
		},
	)

	log.Initialize(ctx)

	alphavantageClient := &httpclient.RequestClient{
		Identifier: httpclient.AlphaVantage,
		Host:       "www.alphavantage.co",
		Scheme:     "https",
		Authority:  "www.alphavantage.co",
	}

	ev := env.NewEnv(
		env.WithAlphaVantageHttpConn(alphavantageClient),
	)

	r := router.SetupRouter(ctx, ev, nrApp)
	srv := &http.Server{
		Addr:         "0.0.0.0:" + os.Getenv("PORT"),
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.ErrorfWithContext(ctx, "unable to start http server")
	}
}
