package server

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config *Config
	engine *gin.Engine
}

func NewServer(config *Config) *Server {
	return &Server{
		config: config,
		engine: gin.Default(),
	}
}

func (s *Server) generateConnectionString() string {
	if s.config.Port == "" {
		s.config.Port = "8080"
	}

	return fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
}

// Run starts the server and listens for incoming requests on the specified address.
func (s *Server) Run() error {
	corsConfig := s.corsConfig()
	s.engine.Use(corsConfig)

	s.healthCheck()

	return s.engine.Run(s.generateConnectionString())
}

func (s *Server) corsConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins: false,
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Authorization", "content-type"},
		AllowHeaders:     []string{"Authorization", "content-type "},
	})
}

func (s *Server) healthCheck() {
	s.engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
