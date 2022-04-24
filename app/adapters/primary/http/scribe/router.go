package scribe

import (
	"github.com/christopher.hachey/scribe/app/domain/scribe/usecase"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

// Router calls the various handlers, and has a copy of the interactor so that it can call the various use cases
type Router struct {
	logger         *zerolog.Logger
	useCaseHandler usecase.ScribePrimaryPorts
}

func NewRouter(uch usecase.ScribePrimaryPorts, logger *zerolog.Logger) Router {
	return Router{
		logger:         logger,
		useCaseHandler: uch,
	}
}

func (r Router) BindGinRoutes(e *gin.Engine) {
	api := e.Group("/api")

	r.textRoutes(api)
}

func (r Router) textRoutes(api *gin.RouterGroup) {
	text := api.Group("/text")
	text.GET("/id/:textId", r.getText)
}
