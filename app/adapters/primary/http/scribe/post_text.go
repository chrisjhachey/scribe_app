package scribe

import (
	"github.com/gin-gonic/gin"
)

func (r Router) postText(c *gin.Context) {
	r.logger.Info().Msg("called primary adapter for http postText!")

	r.useCaseHandler.PostText()
}
