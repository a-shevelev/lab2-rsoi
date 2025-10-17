package server

import (
	"fmt"
	handlers "lab2-rsoi/library-system/internal/handlers/http/v1"
	"lab2-rsoi/library-system/internal/repo"
	"lab2-rsoi/library-system/internal/service"
	"lab2-rsoi/library-system/pkg/postgres"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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

	if err := s.InitDocsRoutes(); err != nil {
		log.Info("Docs routes initialization failed")
	}

	v1 := s.GinRouter.Group("/api/v1")

	//personRepo := repo.New(s.DB)
	library := repo.NewLibraryRepo(s.DB)
	//ctx := context.Background()
	//log.Info(library.FetchLibrariesByCity(ctx, "Москва", 1, 1))
	//log.Info(library.FetchBooksByLibrary(ctx, "83575e12-7ce0-48ee-9931-51919ff3c9ee", true, 1, 1))
	//log.Info(library.IncreaseCount(ctx, 1, 1))
	libraryService := service.NewLibraryService(library)
	libraryHandler := handlers.New(libraryService)
	//personService := service.New(personRepo)
	//personHandler := handlers.New(personService)
	//personHandler.RegisterRoutes(v1)
	libraryHandler.RegisterRoutes(v1)

	return nil
}

func (s *Server) Run() error {
	return s.GinRouter.Run(fmt.Sprintf("%s:%d", s.Host, s.Port))
}
