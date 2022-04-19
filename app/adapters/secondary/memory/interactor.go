package memory

import usecase "github.com/christopher.hachey/scribe/app/domain/scribe/usecase"

type ScribeSecondaryInteractor struct {
	m map[string]string
}

func New() usecase.TextPersistanceSecondaryPort {
	ssi := ScribeSecondaryInteractor{
		make(map[string]string),
	}

	return ssi
}
