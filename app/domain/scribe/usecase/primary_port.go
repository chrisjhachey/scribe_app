package usecase

type ScribePrimaryPort interface {
	TextPrimaryPort
}

type TextPrimaryPort interface {
	GetText()
}
