package adapters

import (
	"errors"
	"sync"

	"github.com/christopher.hachey/scribe/internal/scribe/domain/scribe"
	"github.com/rs/zerolog"
)

type ScribeMemoryAdapter struct {
	memoryMap map[string]interface{}
	lock      *sync.RWMutex
	logger    *zerolog.Logger
}

func NewScribeMemoryAdapter(logger *zerolog.Logger) *ScribeMemoryAdapter {
	m := make(map[string]interface{})

	return &ScribeMemoryAdapter{
		memoryMap: m,
		lock:      &sync.RWMutex{},
		logger:    logger,
	}
}

func (sma ScribeMemoryAdapter) GetText(uri string) (*scribe.Text, error) {
	sma.logger.Info().Msg("successfully called get text adapter!")

	sma.lock.RLock() // blocks for reading
	defer sma.lock.RUnlock()

	record, ok := sma.memoryMap[uri]

	if !ok {
		return &scribe.Text{}, errors.New("Could not find the text")
	}

	text, ok := record.(scribe.Text)

	if !ok {
		return &scribe.Text{}, errors.New("Trouble deserializing to type Text")
	}

	return &text, nil

}

func (sma ScribeMemoryAdapter) PostText(text *scribe.Text) (*scribe.Text, error) {
	sma.logger.Info().Msg("successfully called post text adapter!")

	sma.lock.Lock() // blocks for writing
	defer sma.lock.Unlock()

	sma.memoryMap[text.URI] = *text

	newText, ok := sma.memoryMap[text.URI]

	if !ok {
		return nil, errors.New("Failed to insert the Text")
	}

	t, ok := newText.(scribe.Text)

	if !ok {
		return &scribe.Text{}, errors.New("Trouble deserializing to type Text")
	}

	return &t, nil
}
