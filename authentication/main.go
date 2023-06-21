package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nfwGytautas/wdtk-go-backend/microservice"
	"github.com/nfwGytautas/wdtk-services/auth/endpoints"
)

func main() {
	log.Println("Running WDTK authentication service")

	config, err := microservice.ReadConfig()
	if err != nil {
		panic(err)
	}

	err = setupAuthService(config)
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.SetTrustedProxies([]string{
		config.GatewayIp,
	})

	public := r.Group("/")
	public.POST("Login", endpoints.Login)
	public.POST("Register", endpoints.Register)

	private := r.Group("/", microservice.AuthenticationMiddleware())
	private.GET("Me", endpoints.Me)

	r.Run(config.RunAddress)
}
