package scribe

import (
	"github.com/christopher.hachey/scribe/app/domain/scribe/usecase"
	"github.com/gin-gonic/gin"
)

type RouterHandler struct {
	useCaseHandler usecase.ScribePrimaryPort
}

func NewRouterHandler(uch usecase.ScribePrimaryPort) RouterHandler {
	return RouterHandler{
		useCaseHandler: uch,
	}
}

func (rh RouterHandler) BindGinRoutes(e *gin.Engine) {
	api := e.Group("/api")

	rh.textRoutes(api)
}

func (rh RouterHandler) textRoutes(api *gin.RouterGroup) {
	text := api.Group("/text")
	text.GET("/id/:textId", rh.getText)
}
