package usecase

import (
	"github.com/christopher.hachey/scribe/app/adapters/primary/http/scribe/deserializer"
	"github.com/christopher.hachey/scribe/app/domain/scribe/model"
	"github.com/google/uuid"
)

func (i ScribePrimaryInteractor) PostText(text deserializer.TextPostRequest) (model.Text, error) {
	i.Logger.Info().Msg("Called post text in the use case!")

	uri := *text.Name + uuid.New().String()
	t := model.Text{
		Name:   text.Name,
		Author: text.Author,
		Uri:    &uri,
	}

	newText, err := i.Storage.ScribeSecondaryPorts.PostText(t)

	if err != nil {
		return model.Text{}, err
	}

	return *newText, nil
}
