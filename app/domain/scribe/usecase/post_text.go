package usecase

func (i ScribePrimaryInteractor) PostText() {
	i.Logger.Info().Msg("Called post text in the use case!")

	i.Storage.ScribeSecondaryPorts.PostText()
}
