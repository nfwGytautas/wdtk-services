package main

import (
	"log"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/nfwGytautas/wdtk-go-backend/microservice"
)

func main() {
	log.Println("Running WDTK http server")

	config, err := microservice.ReadConfig()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./", true)))
	r.Run(config.RunAddress)
}
