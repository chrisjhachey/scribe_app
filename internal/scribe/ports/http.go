package ports

import (
	"net/http"

	"github.com/christopher.hachey/scribe/internal/scribe/app"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type RouterHandler struct {
	service app.ScribeService
	logger  *zerolog.Logger
}

func NewRouterHandler(service app.ScribeService, logger *zerolog.Logger) RouterHandler {
	return RouterHandler{
		service: service,
		logger:  logger,
	}
}

func (rh RouterHandler) GetText(c *gin.Context, uri string) {
	rh.logger.Info().Msg("successfully called get text router handler")

	text, err := rh.service.GetText(uri)

	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}

	resp := Text{
		NewText: NewText{
			Name:   text.Name,
			Author: &text.Author,
		},
		Uri: text.URI,
	}

	c.JSON(http.StatusOK, resp)
}

func (rh RouterHandler) PostText(c *gin.Context) {
	rh.logger.Info().Msg("successfully called post text router handler")

	nt := NewText{}
	err := c.BindJSON(&nt)

	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}

	text, err := rh.service.PostText(nt.Name, *nt.Author)

	resp := Text{
		NewText: nt,
		Uri:     text.URI,
	}

	c.JSON(http.StatusOK, resp)
}
