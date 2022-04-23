package server

import (
	factory "github.com/christopher.hachey/scribe/app/infrastructure/factory"
	"github.com/spf13/viper"
)

type Server struct {
	UseCaseHandlerFactory factory.UseCaseHandlerFactory
	GinServer             *GinServer
	DatabaseURL           string
	// Logger *logger.StandardLogger
	// Exiter exiter.Exiter
}

func (s Server) Initialize() {
	primaryPorts := viper.GetStringSlice("PRIMARY_PORTS")

	if len(primaryPorts) > 0 {
		for _, v := range primaryPorts {
			// If the http primary port is enabled, we inject the gin server and bind the scribe primary adapters
			if v == "http" {
				s.UseCaseHandlerFactory = factory.UseCaseHandlerFactory{}

				if s.GinServer == nil {
					s.GinServer = NewServer(viper.GetInt("PORT"))
					s.GinServer.BindScribePrimaryAdaptersToGinServer(s.UseCaseHandlerFactory)
					s.GinServer.Start()
				}
			} else {
				return
			}
		}
	}
}
