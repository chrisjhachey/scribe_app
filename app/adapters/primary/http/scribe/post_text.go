package scribe

import (
	"net/http"

	"github.com/christopher.hachey/scribe/app/adapters/primary/http/scribe/deserializer"
	"github.com/gin-gonic/gin"
)

func (r Router) postText(c *gin.Context) {
	r.logger.Info().Msg("called primary adapter for http postText!")

	tpr := deserializer.TextPostRequest{}

	err := c.BindJSON(&tpr)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	text, err := r.useCaseHandler.PostText(tpr)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, text)
}
