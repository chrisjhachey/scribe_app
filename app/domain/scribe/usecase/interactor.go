package usecase

import "github.com/rs/zerolog"

type ScribePrimaryInteractor struct {
	Logger  *zerolog.Logger
	Storage ScribeSecondaryAdapters
}

type ScribeSecondaryAdapters struct {
	ScribeSecondaryPorts ScribeSecondaryPorts
}
