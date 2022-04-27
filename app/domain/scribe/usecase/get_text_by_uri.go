package usecase

import "github.com/christopher.hachey/scribe/app/domain/scribe/model"

func (i ScribePrimaryInteractor) GetText(textURI string) (model.Text, error) {
	i.Logger.Info().Msg("Called get text in the use case!")

	text, err := i.Storage.ScribeSecondaryPorts.GetText(textURI)

	if err != nil {
		return text, err
	}

	return text, nil
}
