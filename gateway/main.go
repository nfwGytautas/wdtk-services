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
	httpsEnabled, httpsDefined := config.UserDefines["HTTPS"]
	if httpsDefined && httpsEnabled.(bool) {
		// HTTPS requires key and certification file
		certFile, certDefined := config.UserDefines["CertFile"]
		keyFile, keyDefined := config.UserDefines["KeyFile"]

		if !certDefined || !keyDefined {
			log.Fatal("'CertFile' and 'KeyFile' must be defined")
			return
		}

		r.RunTLS(config.RunAddress, certFile.(string), keyFile.(string))
	} else {
		// HTTP
		r.Run(config.RunAddress)
	}
}
