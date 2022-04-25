package memory

func (si ScribeMemorySecondaryInteractor) PostText() {
	si.logger.Info().Msg("Just posted text to repo!")

	// mm := *si.memoryMap
	// text := mm[textURI]
}
