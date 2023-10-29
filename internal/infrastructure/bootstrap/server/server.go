package server

import (
	"fmt"

	tweetservice "github.com/bypepe77/Twitter-Clone-for-Golang/internal/application/tweet"
	service "github.com/bypepe77/Twitter-Clone-for-Golang/internal/application/user"
	"github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/api/auth"
	"github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/api/middlewares"
	tweetapi "github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/api/tweet"
	jwtManager "github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/jwt"
	repositories "github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/repositories/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.temporal.io/sdk/client"
	"gorm.io/gorm"
)

type Server struct {
	config         *Config
	engine         *gin.Engine
	db             *gorm.DB
	temporalClient client.Client
}

func NewServer(config *Config, db *gorm.DB, temporalClient client.Client) *Server {
	return &Server{
		config:         config,
		engine:         gin.Default(),
		db:             db,
		temporalClient: temporalClient,
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
	jwtManager := jwtManager.New("secret", "twitter-clone")
	userRepository := repositories.NewUserRepository(s.db)
	userService := service.NewUserService(userRepository, jwtManager)
	authAPI := auth.New(userService)
	authRouter := auth.NewRouter(*s.engine.Group("/auth"), authAPI)
	authRouter.Register()

	//Register tweet routes
	tweetService := tweetservice.New(s.temporalClient)
	tweetAPI := tweetapi.New(tweetService, jwtManager)
	tweetRouter := tweetapi.NewRouter(*s.engine.Group("/tweet", middlewares.Authorize(jwtManager)), tweetAPI)
	tweetRouter.Register()

}
