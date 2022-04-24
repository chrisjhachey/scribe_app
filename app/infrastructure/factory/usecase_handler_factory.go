package factory

import (
	"github.com/christopher.hachey/scribe/app/adapters/secondary/memory"
	"github.com/christopher.hachey/scribe/app/domain/scribe/usecase"
	"github.com/spf13/viper"
)

type UseCaseHandlerFactory struct {
}

func NewUseCaseHandlerFactory() UseCaseHandlerFactory {
	return UseCaseHandlerFactory{}
}

func (uchf UseCaseHandlerFactory) BuildScribeUseCaseHandler() usecase.ScribePrimaryPorts {
	var useCaseHandler usecase.ScribePrimaryPorts
	sa := viper.GetString("SECONDARY_ADAPTERS")

	switch sa {
	case "memory":
		useCaseHandler = uchf.buildScribeMemoryUseCaseHandler()
	}

	return useCaseHandler
}

func (uchf UseCaseHandlerFactory) buildScribeMemoryUseCaseHandler() usecase.ScribePrimaryInteractor {
	return usecase.ScribePrimaryInteractor{
		Storage: usecase.ScribeSecondaryAdapters{
			ScribeSecondaryPorts: memory.New(),
		},
	}
}
