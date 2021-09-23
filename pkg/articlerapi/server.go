package articlerapi

import (
	"context"
	"github.com/Nealoth/articler-api/pkg/articlerapi/configuration"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Conf   configuration.Configuration
	Engine *gin.Engine
}

func NewServer(conf configuration.Configuration) *Server {
	s := &Server{
		Conf:   conf,
		Engine: gin.Default(),
	}

	return s
}

func (s *Server) Start(ctx context.Context) error {
	err := s.Engine.Run(":" + s.Conf.ServerConfiguration.Port)

	if err != nil {
		return err
	}

	return nil
}
