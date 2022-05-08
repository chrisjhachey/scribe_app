package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/christopher.hachey/scribe/internal/common/config"
	"github.com/christopher.hachey/scribe/internal/common/server"
	"github.com/christopher.hachey/scribe/internal/scribe/adapters"
	"github.com/christopher.hachey/scribe/internal/scribe/app"
	"github.com/christopher.hachey/scribe/internal/scribe/domain/scribe"
	"github.com/christopher.hachey/scribe/internal/scribe/ports"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	config.CobraInitialization()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())

	defer func() {
		signal.Stop(sig)
		cancel()
	}()

	// Inject adapters
	var repo scribe.Repository

	switch viper.GetString("SECONDARY_ADAPTERS") {
	case "memory":
		repo = adapters.NewScribeMemoryAdapter()
	case "pgsql":

	}

	scribe := app.NewScribeService(repo)

	rh := ports.NewRouterHandler(scribe)

	ginServer := server.NewServer(viper.GetInt("PORT"))

	ports.RegisterHandlersWithOptions(ginServer.Engine, rh, ports.GinServerOptions{
		BaseURL: "/api",
		Middlewares: []ports.MiddlewareFunc{
			func(c *gin.Context) {
				fmt.Println("I'm Middleware 1!")
			},
			func(c *gin.Context) {
				fmt.Println("I'm Middleware 2!")
			},
		},
	})

	ginServer.Start()

	select {
	case <-sig:
		cancel()
		os.Exit(1)
	case <-ctx.Done():
	}
}
