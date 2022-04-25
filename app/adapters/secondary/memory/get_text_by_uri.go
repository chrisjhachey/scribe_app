package memory

func (si ScribeMemorySecondaryInteractor) GetText(textURI string) {
	si.logger.Info().Msg("Just got text from repo!")

	// mm := *si.memoryMap
	// text := mm[textURI]
}
