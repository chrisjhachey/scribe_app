package memory

import (
	"encoding/json"

	"github.com/christopher.hachey/scribe/app/domain/scribe/model"
)

func (si ScribeMemorySecondaryInteractor) GetText(textURI string) (model.Text, error) {
	si.logger.Info().Msg("Just got text from repo!")

	mm := *si.memoryMap
	text := mm[textURI]

	var textStruct model.Text
	err := json.Unmarshal([]byte(text), &textStruct)

	if err != nil {
		return model.Text{}, err
	}

	return textStruct, nil
}
