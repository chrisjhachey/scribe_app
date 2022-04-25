package usecase

type ScribePrimaryPorts interface {
	TextPrimaryPort
}

type TextPrimaryPort interface {
	GetText(textURI string)
	PostText()
}
