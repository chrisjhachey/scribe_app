package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/christopher.hachey/scribe/internal/common/config"
	"github.com/christopher.hachey/scribe/internal/common/server"
	"github.com/christopher.hachey/scribe/internal/scribe/adapters"
	"github.com/christopher.hachey/scribe/internal/scribe/app"
	"github.com/christopher.hachey/scribe/internal/scribe/domain/scribe"
	"github.com/christopher.hachey/scribe/internal/scribe/ports"
	"github.com/rs/zerolog"
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

	writer := zerolog.ConsoleWriter{Out: os.Stderr}
	logger := zerolog.New(writer).With().Timestamp().Logger()

	// Inject adapters
	var repo scribe.Repository

	switch viper.GetString("SECONDARY_ADAPTERS") {
	case "memory":
		repo = adapters.NewScribeMemoryAdapter(&logger)
	case "pgsql":

	}

	scribe := app.NewScribeService(repo, &logger)

	rh := ports.NewRouterHandler(scribe, &logger)

	ginServer := server.NewServer(viper.GetInt("PORT"))

	ports.RegisterHandlersWithOptions(ginServer.Engine, rh, ports.GinServerOptions{
		BaseURL:     "/api",
		Middlewares: []ports.MiddlewareFunc{
			// func(c *gin.Context) {
			// 	fmt.Println("I'm Middleware 1!")
			// }
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
