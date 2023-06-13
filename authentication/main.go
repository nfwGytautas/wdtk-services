package main

import (
	"log"

	"github.com/nfwGytautas/wdtk-go-backend/microservice"
)

func main() {
	log.Println("Running WDTK authentication service")

	if err := microservice.RegisterService[AuthData](microservice.ServiceDescription[AuthData]{
		SetupFn: setupAuthService,
	}, []microservice.ServiceEndpoint{
		{
			Type:            microservice.ENDPOINT_TYPE_POST,
			Name:            "Login/",
			Fn:              loginEndpoint,
			EndpointContext: nil,
		},
	}); err != nil {
		log.Println(err)
		panic("Failed to register authentication service")
	}
}
