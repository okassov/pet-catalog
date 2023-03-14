package main

import (
  "log"

  "github.com/okassov/pet-catalog/internal/app"
  "github.com/okassov/pet-catalog/config"
)

func main() {

  config, err := config.LoadConfig(".")
  if err != nil {
    log.Fatal("Cannot load config: ", err)
  }
  app.Run(config)
}
