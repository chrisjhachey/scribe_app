package server

import (
	"os"

	factory "github.com/christopher.hachey/scribe/app/infrastructure/factory"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type Server struct {
	UseCaseHandlerFactory factory.UseCaseHandlerFactory
	GinServer             *GinServer
	DatabaseURL           string
	Logger                *zerolog.Logger
	// Exiter exiter.Exiter
}

func (s Server) Initialize() {
	writer := zerolog.ConsoleWriter{Out: os.Stderr}
	logger := zerolog.New(writer).With().Timestamp().Logger()
	s.Logger = &logger

	primaryPorts := viper.GetStringSlice("PRIMARY_PORTS")

	if len(primaryPorts) > 0 {
		for _, v := range primaryPorts {
			// If the http primary port is enabled, we inject the gin server and bind the scribe primary adapters
			if v == "http" {
				s.UseCaseHandlerFactory = factory.UseCaseHandlerFactory{
					Logger: s.Logger,
				}

				if s.GinServer == nil {
					s.GinServer = NewServer(viper.GetInt("PORT"), s.Logger)
					s.GinServer.BindScribePrimaryAdaptersToGinServer(s.UseCaseHandlerFactory)
					s.GinServer.Start()
				}
			} else {
				return
			}
		}
	}
}
