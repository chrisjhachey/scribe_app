package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/christopher.hachey/scribe/app/config"
	"github.com/christopher.hachey/scribe/app/infrastructure/server"
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

	asc := server.ServerConstructor{}
	asc.Initialize()

	select {
	case <-sig:
		cancel()
		os.Exit(1)
	case <-ctx.Done():
	}
}
