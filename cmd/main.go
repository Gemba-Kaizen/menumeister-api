package main

import (
	"log"

	"github.com/Gemba-Kaizen/menumeister-api/config"
	"github.com/Gemba-Kaizen/menumeister-api/pkg/auth"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {

	// Load config
	c, err := config.LoadConfig()

	if err!= nil {
    log.Fatalln("Failed at config", err)
  }

	// Gin Engine
	r := gin.Default()

	// Enable CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{c.MenuMeisterFronEnd}
	r.Use(cors.New(config))

	// get registered routes function from ping package
	auth.RegisterRoutes(r, &c)

	r.Run(c.Port)
}