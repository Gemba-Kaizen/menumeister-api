package main

import (
	"log"

	"github.com/Gemba-Kaizen/menumeister-api/config"
	"github.com/Gemba-Kaizen/menumeister-api/pkg/ping"
	"github.com/gin-gonic/gin"
)

func main() {

	// Load config
	c, err := config.LoadConfig()

	if err!= nil {
    log.Fatalln("Failed at config", err)
  }

	// Gin Engine
	r := gin.Default()

	// get registered routes function from ping package
	ping.RegisterRoutes(r, &c)

	r.Run(c.Port)
}