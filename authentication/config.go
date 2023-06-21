package main

import (
	"log"

	"github.com/nfwGytautas/wdtk-go-backend/microservice"
	"github.com/nfwGytautas/wdtk-services/auth/context"
)

func setupAuthService(config *microservice.MicroserviceConfig) error {
	log.Println("Setting up authentication service")

	log.Printf("- RunAddress       : %s", config.RunAddress)
	log.Printf("- Gateway          : %s", config.GatewayIp)
	log.Printf("- ConnectionString : %s", config.UserDefines["connectionString"].(string))
	log.Printf("- ApiKey           : %s", config.ApiKey)

	err := context.Context.SetupDatabase(config.UserDefines["connectionString"].(string))
	return err
}
