package app

import (
	"fmt"

	"github.com/christopher.hachey/scribe/internal/scribe/domain/scribe"
)

type ScribeService struct {
	repository scribe.Repository
}

func NewScribeService(repository scribe.Repository) ScribeService {
	return ScribeService{
		repository: repository,
	}
}

func (ss ScribeService) GetText(uri string) (*scribe.Text, error) {
	fmt.Println("successfully called get text use case!")
	text, err := ss.repository.GetText(uri)

	return text, err
}

func (ss ScribeService) PostText(name string, author string) (*scribe.Text, error) {
	fmt.Println("successfully called post text use case!")
	uri := fmt.Sprintf("sws::text::%s", name)

	t := scribe.Text{
		Name:   name,
		Author: author,
		URI:    uri,
	}

	text, err := ss.repository.PostText(&t)

	return text, err
}
