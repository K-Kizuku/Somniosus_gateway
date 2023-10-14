package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/K-Kizuku/Somniosus_gateway/pkg/lib/config"
	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.NewAppConfig()
	e := echo.New()
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              cfg.Host,
		TracesSampleRate: 1.0,
	}); err != nil {
		log.Fatalf("Sentry initialization failed: %v\n", err)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(sentryecho.New(sentryecho.Options{}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/panic", func(c echo.Context) error {
		panic("kizuku")
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
