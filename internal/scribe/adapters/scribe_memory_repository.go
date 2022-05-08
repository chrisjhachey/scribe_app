package adapters

import (
	"errors"
	"fmt"
	"sync"

	"github.com/christopher.hachey/scribe/internal/scribe/domain/scribe"
)

type ScribeMemoryAdapter struct {
	memoryMap map[string]interface{}
	lock      *sync.RWMutex
}

func NewScribeMemoryAdapter() *ScribeMemoryAdapter {
	m := make(map[string]interface{})

	return &ScribeMemoryAdapter{
		memoryMap: m,
		lock:      &sync.RWMutex{},
	}
}

func (sma ScribeMemoryAdapter) GetText(uri string) (*scribe.Text, error) {
	sma.lock.RLock() // blocks for reading
	defer sma.lock.RUnlock()

	fmt.Println("successfully called get text adapter!")

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
	sma.lock.Lock() // blocks for writing
	defer sma.lock.Unlock()

	fmt.Println("successfully called post text adapter!")

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
