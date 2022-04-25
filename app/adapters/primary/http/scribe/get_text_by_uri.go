package scribe

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (r Router) getText(c *gin.Context) {
	uri := c.Params.ByName("textURI")
	r.logger.Info().Msg(fmt.Sprintf("called primary adapter for http getText with uri %s", uri))

	r.useCaseHandler.GetText(uri)
}
