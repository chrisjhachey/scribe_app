package scribe

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r Router) getText(c *gin.Context) {
	uri := c.Params.ByName("textURI")
	r.logger.Info().Msg(fmt.Sprintf("called primary adapter for http getText with uri %s", uri))

	text, err := r.useCaseHandler.GetText(uri)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, text)
}
