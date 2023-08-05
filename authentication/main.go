package main

import (
	"fmt"
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

	setupRegisterEndpoint(r, config)

	private := r.Group("/", microservice.AuthenticationMiddleware())
	private.GET("Me", endpoints.Me)

	r.Run(config.RunAddress)
}

func setupRegisterEndpoint(r *gin.Engine, config *microservice.MicroserviceConfig) {
	if !config.UserDefines["allowRegistration"].(bool) {
		return
	}

	println()
	defer println()

	group := r.Group("/")
	values, exists := config.UserDefines["registerRoles"]
	if exists && len(values.([]interface{})) > 0 {
		log.Println("Allowing registration for:")
		roles := make([]string, len(values.([]interface{})))
		for i, v := range values.([]interface{}) {
			roles[i] = fmt.Sprint(v)

			if roles[i] == "" {
				log.Panicf("Empty role at index %v not allowed", i)
			}
			log.Printf("\t - %s", roles[i])
		}

		// Private
		group.Use(microservice.AuthorizationMiddleware(roles))
	}
	group.POST("Register", endpoints.Register)
}
