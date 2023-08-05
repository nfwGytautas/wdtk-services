package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nfwGytautas/wdtk-go-backend/microservice"
)

func main() {
	log.Println("Running WDTK API gateway")

	// Read generated config
	config, err := microservice.ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	// Create gin engine
	r := gin.Default()
	// r.SetTrustedProxies(nil)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	r.Use(cors.New(corsConfig))

	// Configure forwarding routes
	SetupRoutes(config, r)

	// Run gin and block routine
	r.Run(config.RunAddress)
}
