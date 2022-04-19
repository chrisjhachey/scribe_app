package server

import (
	factory "github.com/christopher.hachey/scribe/app/infrastructure/factory"
	"github.com/spf13/viper"
)

type ServerConstructor struct {
	UseCaseHandlerFactory factory.UseCaseHandlerFactory
	GinServer             *GinServer
	DatabaseURL           string
	// Logger *logger.StandardLogger
	// Exiter exiter.Exiter
}

func (sc ServerConstructor) Initialize() {
	primaryPorts := viper.GetStringSlice("PRIMARY_PORTS")

	if len(primaryPorts) > 0 {
		for _, v := range primaryPorts {
			if v == "gin" {
				sc.UseCaseHandlerFactory = factory.NewUseCaseHandlerFactory()

				if sc.GinServer == nil {
					sc.GinServer = NewServer(viper.GetInt("PORT"))
					sc.GinServer.BindScribePrimaryAdaptersToGinServer(sc.UseCaseHandlerFactory)
					sc.GinServer.Start()
				}
			} else {
				return
			}
		}
	}
}
