package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/okassov/pet-catalog/config"
	v1 "github.com/okassov/pet-catalog/internal/controller/http/v1"
	"github.com/okassov/pet-catalog/internal/entity"
	"github.com/okassov/pet-catalog/internal/usecase"
	"github.com/okassov/pet-catalog/internal/usecase/repository"
	"github.com/okassov/pet-catalog/pkg/httpserver"
	"github.com/okassov/pet-catalog/pkg/logger"
	"github.com/okassov/pet-catalog/pkg/postgres"
)

func Run(config config.Config) {

	// Init Logger
	l := logger.New()

	// Repository
	pgConnString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.PG.PGUser,
		config.PG.PGPassword,
		config.PG.PGUrl,
		config.PG.PGPort,
		config.PG.PGDatabase)

	pg, err := postgres.New(pgConnString)
	if err != nil {
		l.Error("app - Run - postgres.New: %w", err)
	}
	// l.Info("Successfuly connected to the Database")
	pg.AutoMigrate(&entity.Product{})

	// Use case
	productUseCase := usecase.NewProductUseCase(
		repository.NewProductRepo(*pg),
	)

	handler := gin.New()

	// Routers
	v1.NewRouter(handler, productUseCase, l)

	// Server
	httpServer := httpserver.New(handler, httpserver.Port(config.Server.Port))

	// Shutdown Server
	err = httpServer.Shutdown()
	if err != nil {
		l.Error("app - Run - httpServer.Shutdown: ", err)
	}
}
