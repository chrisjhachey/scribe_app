package usecase

type ScribeSecondaryPorts interface {
	TextPersistanceSecondaryPort
}

type TextPersistanceSecondaryPort interface {
	GetText(textURI string)
	PostText()
}
