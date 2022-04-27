package deserializer

type TextPostRequest struct {
	Name   *string `json:"name" binding:"required"`
	Author *string `json:"author"`
}
