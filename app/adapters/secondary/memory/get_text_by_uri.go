package memory

func (si ScribeMemorySecondaryInteractor) GetText() {
	si.logger.Info().Msg("Just got text from repo!")
}
