package memory

import (
	"encoding/json"

	"github.com/christopher.hachey/scribe/app/domain/scribe/model"
)

func (si ScribeMemorySecondaryInteractor) PostText(text model.Text) (*model.Text, error) {
	si.logger.Info().Msg("Just posted text to repo!")

	b, err := json.Marshal(text)

	if err != nil {
		return nil, err
	}

	mm := *si.memoryMap
	mm[*text.Uri] = string(b)

	return &text, nil
}
