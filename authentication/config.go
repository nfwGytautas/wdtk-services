package main

import (
	"log"
)

func setupAuthService(context *AuthData, config map[string]interface{}) error {
	log.Println("Setting up authentication service")

	log.Printf("- RunAddress       : %s", config["runAddress"])
	log.Printf("- Gateway          : %s", config["gateway"])
	log.Printf("- ConnectionString : %s", config["connectionString"])
	return nil
}
