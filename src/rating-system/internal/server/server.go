package server

import (
	"fmt"

	"rating-system/internal/handlers/http/v1"
	"rating-system/internal/repo"
	"rating-system/internal/service"
	"rating-system/pkg/postgres"

	"github.com/gin-gonic/gin"
	//log "github.com/sirupsen/logrus"
	//log "github.com/sirupsen/logrus"
)

type Server struct {
	Host      string `envconfig:"HOST" required:"true"`
	Port      int    `envconfig:"PORT" required:"true"`
	DB        postgres.Client
	GinRouter *gin.Engine
}

func New(dbc postgres.Client, host string, port int) (*Server, error) {
	s := &Server{
		Host:      host,
		Port:      port,
		DB:        dbc,
		GinRouter: gin.Default(),
	}

	if err := s.initRoutes(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Server) initRoutes() error {
	s.GinRouter.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "pong"})
	})

	v1 := s.GinRouter.Group("/api/v1")

	rateRepo := repo.NewRatingRepo(s.DB)
	//ctx := context.Background()
	//log.Info(library.FetchLibrariesByCity(ctx, "Москва", 1, 1))
	//log.Info(library.FetchBooksByLibrary(ctx, "83575e12-7ce0-48ee-9931-51919ff3c9ee", true, 1, 1))
	//log.Info(library.IncreaseCount(ctx, 1, 1))

	rateService := service.NewRatingService(rateRepo)
	//libraryService := service.NewLibraryService(library)
	//libraryHandler := handlers.New(libraryService)
	//libraryHandler.RegisterRoutes(v1)
	rateHandler := handlers.New(rateService)
	rateHandler.RegisterRoutes(v1)

	return nil
}

func (s *Server) Run() error {
	return s.GinRouter.Run(fmt.Sprintf("%s:%d", s.Host, s.Port))
}
