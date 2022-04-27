package usecase

import (
	"github.com/christopher.hachey/scribe/app/domain/scribe/model"
)

type ScribeSecondaryPorts interface {
	TextPersistanceSecondaryPort
}

type TextPersistanceSecondaryPort interface {
	GetText(textURI string) (model.Text, error)
	PostText(model.Text) (*model.Text, error)
}
