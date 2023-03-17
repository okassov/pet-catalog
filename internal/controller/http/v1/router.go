package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/okassov/pet-catalog/internal/usecase"
	"github.com/okassov/pet-catalog/pkg/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
	// Swagger docs.
	// _ "github.com/okassov/pet-catalog/docs"
)

//	@title			Pet Catalog Service
//	@description	Using a catalog service
//	@version		1.0
//	@host			localhost:8080
//	@BasePath		/v1
func NewRouter(handler *gin.Engine, c usecase.Catalog, l logger.LoggerInterface) {

	// authMiddleware := NewAuthMiddleware(a)

	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// // Added/Skip OTLP Middleware from infra routes
	// handler.Use(SkipOTLPMiddleware)

	// // Swagger
	// swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	// handler.GET("/swagger/*any", swaggerHandler)

	// Healthcheck endpoint
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", func(c *gin.Context) { promhttp.Handler().ServeHTTP(c.Writer, c.Request) })

	// Routers
	h := handler.Group("/v1")
	{
		newProductRoutes(h, c, l)
		newCategoryRoutes(h, c, l)
	}

}
