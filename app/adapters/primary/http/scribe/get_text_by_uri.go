package scribe

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (rh RouterHandler) getText(c *gin.Context) {
	fmt.Println("called primary adapter for http getText!")

	rh.useCaseHandler.GetText()
}
