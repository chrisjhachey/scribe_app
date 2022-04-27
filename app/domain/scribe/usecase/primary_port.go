package usecase

import (
	"github.com/christopher.hachey/scribe/app/adapters/primary/http/scribe/deserializer"
	"github.com/christopher.hachey/scribe/app/domain/scribe/model"
)

type ScribePrimaryPorts interface {
	TextPrimaryPort
}

type TextPrimaryPort interface {
	GetText(textURI string) (model.Text, error)
	PostText(text deserializer.TextPostRequest) (model.Text, error)
}
