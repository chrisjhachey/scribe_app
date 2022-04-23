package server

import (
	"crypto/tls"
	"net"
	"net/http"
	"strconv"

	"fmt"

	"github.com/christopher.hachey/scribe/app/adapters/primary/http/scribe"
	factory "github.com/christopher.hachey/scribe/app/infrastructure/factory"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type GinServer struct {
	Server *http.Server
	Engine *gin.Engine
}

func NewServer(port int) *GinServer {
	gin.SetMode(viper.GetString("GIN_MODE"))
	router := gin.New()

	addr := ":" + strconv.Itoa(port)

	return &GinServer{
		Engine: router,
		Server: &http.Server{
			Addr:    addr,
			Handler: router,
			TLSConfig: &tls.Config{
				MinVersion: tls.VersionTLS13,
			},
		},
	}
}

func (s *GinServer) BindScribePrimaryAdaptersToGinServer(uchf factory.UseCaseHandlerFactory) {
	scribeInteractor := uchf.BuildScribeUseCaseHandler()

	scribe.NewRouter(scribeInteractor).BindGinRoutes(s.Server.Handler.(*gin.Engine))
}

func (s *GinServer) Start() {
	fmt.Printf("Now Listening on port %s", s.Server.Addr)

	ln, err := net.Listen("tcp", s.Server.Addr)

	if err != nil {
		fmt.Printf("Error listening on %s", s.Server.Addr)
	}

	go func() {
		defer ln.Close()

		if err := s.Server.Serve(ln); err != nil {
			fmt.Println("Error Serving...")
		}
	}()
}
