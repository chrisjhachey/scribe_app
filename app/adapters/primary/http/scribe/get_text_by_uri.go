package scribe

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (r Router) getText(c *gin.Context) {
	fmt.Println("called primary adapter for http getText!")

	r.useCaseHandler.GetText()
}
