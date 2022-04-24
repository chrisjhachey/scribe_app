package memory

import usecase "github.com/christopher.hachey/scribe/app/domain/scribe/usecase"

type ScribeMemorySecondaryInteractor struct {
	m *map[string]string
}

func New() usecase.TextPersistanceSecondaryPort {
	m := make(map[string]string)
	ssi := ScribeMemorySecondaryInteractor{
		&m,
	}

	return ssi
}
