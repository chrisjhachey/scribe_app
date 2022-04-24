package usecase

func (i ScribePrimaryInteractor) GetText() {
	i.Logger.Info().Msg("Called get text in the use case!")

	i.Storage.ScribeSecondaryPorts.GetText()
}
