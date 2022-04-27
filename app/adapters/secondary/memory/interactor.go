package memory

import (
	usecase "github.com/christopher.hachey/scribe/app/domain/scribe/usecase"
	"github.com/rs/zerolog"
)

type ScribeMemorySecondaryInteractor struct {
	logger    *zerolog.Logger
	memoryMap *map[string]string
}

func New(logger *zerolog.Logger) usecase.TextPersistanceSecondaryPort {
	m := make(map[string]string)
	ssi := ScribeMemorySecondaryInteractor{
		logger,
		&m,
	}

	return ssi
}
