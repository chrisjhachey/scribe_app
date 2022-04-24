package scribe

import (
	"github.com/gin-gonic/gin"
)

func (r Router) getText(c *gin.Context) {
	r.logger.Info().Msg("called primary adapter for http getText!")

	r.useCaseHandler.GetText()
}
