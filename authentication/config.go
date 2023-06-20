package main

import (
	"log"

	"github.com/nfwGytautas/wdtk-services/auth/context"
)

func setupAuthService(context *context.AuthData, config map[string]interface{}) error {
	log.Println("Setting up authentication service")

	log.Printf("- RunAddress       : %s", config["runAddress"])
	log.Printf("- Gateway          : %s", config["gateway"])
	log.Printf("- ConnectionString : %s", config["connectionString"])
	log.Printf("- ApiKey		   : %s", config["apiKey"])

	err := context.SetupDatabase(config["connectionString"].(string))
	return err
}
