package usecase

import (
	"fmt"
)

func (si ScribePrimaryInteractor) GetText() {
	fmt.Println("Called get text in the use case!")

	si.Storage.ScribeSecondaryPorts.GetText()
}
