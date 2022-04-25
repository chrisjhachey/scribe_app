package usecase

func (i ScribePrimaryInteractor) GetText(textURI string) {
	i.Logger.Info().Msg("Called get text in the use case!")

	i.Storage.ScribeSecondaryPorts.GetText(textURI)
}
