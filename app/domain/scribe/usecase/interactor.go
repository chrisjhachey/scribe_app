package usecase

type ScribePrimaryInteractor struct {
	Storage ScribeSecondaryAdapters
}

type ScribeSecondaryAdapters struct {
	TextSecondaryPorts TextPersistanceSecondaryPort
}
