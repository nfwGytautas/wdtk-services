package main

import (
	"log"
)

func setupAuthService(context *AuthData, config map[string]interface{}) error {
	log.Println("Setting up authentication service")

	log.Printf("- RunAddress       : %s", config["RunAddress"])
	log.Printf("- Gateway          : %s", config["Gateway"])
	log.Printf("- ConnectionString : %s", config["ConnectionString"])
	return nil
}
