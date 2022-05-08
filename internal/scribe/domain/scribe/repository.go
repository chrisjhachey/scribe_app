package scribe

type Repository interface {
	GetText(uri string) (*Text, error)
	PostText(text *Text) (*Text, error)
}
