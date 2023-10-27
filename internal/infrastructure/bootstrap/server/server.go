package server

import (
	"fmt"

	service "github.com/bypepe77/Twitter-Clone-for-Golang/internal/application/user"
	"github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/api/auth"
	repositories "github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/repositories/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	config *Config
	engine *gin.Engine
	db     *gorm.DB
}

func NewServer(config *Config, db *gorm.DB) *Server {
	return &Server{
		config: config,
		engine: gin.Default(),
		db:     db,
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
	s.RegisterAuthRoutes()

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

func (s *Server) RegisterAuthRoutes() {
	// TODO: Inject dependencies
	// Register auth routes
	userRepository := repositories.NewUserRepository(s.db)
	userService := service.NewUserService(userRepository)
	authAPI := auth.New(userService)
	authRouter := auth.NewRouter(*s.engine.Group("/auth"), authAPI)
	authRouter.Register()
}
