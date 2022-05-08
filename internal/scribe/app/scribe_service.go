package app

import (
	"fmt"

	"github.com/christopher.hachey/scribe/internal/scribe/domain/scribe"
	"github.com/rs/zerolog"
)

type ScribeService struct {
	repository scribe.Repository
	logger     *zerolog.Logger
}

func NewScribeService(repository scribe.Repository, logger *zerolog.Logger) ScribeService {
	return ScribeService{
		repository: repository,
		logger:     logger,
	}
}

func (ss ScribeService) GetText(uri string) (*scribe.Text, error) {
	ss.logger.Info().Msg("successfully called get text use case!")
	text, err := ss.repository.GetText(uri)

	return text, err
}

func (ss ScribeService) PostText(name string, author string) (*scribe.Text, error) {
	ss.logger.Info().Msg("successfully called post text use case!")

	formattedName, err := scribe.GenerateURIFromName(name)

	if err != nil {
		return &scribe.Text{}, err
	}

	uri := fmt.Sprintf("sws::text::%s", formattedName)

	t := scribe.Text{
		Name:   name,
		Author: author,
		URI:    uri,
	}

	text, err := ss.repository.PostText(&t)

	return text, err
}
