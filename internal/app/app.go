package app

import (
  "github.com/gin-gonic/gin"
  "github.com/okassov/pet-catalog/pkg/httpserver"
  "github.com/okassov/pet-catalog/pkg/logger"
  "github.com/okassov/pet-catalog/config"
)

func Run(config config.Config) {
  
  // Init Logger
  l := logger.New()

  handler := gin.New()

  // Server
  //httpServer := httpserver.New(handler, httpserver.Port(config.Server.port))
  httpServer := httpserver.New(handler, httpserver.Port(":8080"))

  // Shutdown Server
  err := httpServer.Shutdown()
  if err != nil {
    l.Error("app - Run - httpServer.Shutdown: ", err)
  }
}
