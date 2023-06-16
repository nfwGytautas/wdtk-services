package main

import (
	"log"

	"github.com/nfwGytautas/wdtk-go-backend/microservice"
	"github.com/nfwGytautas/wdtk-services/auth/context"
	"github.com/nfwGytautas/wdtk-services/auth/endpoints"
)

func main() {
	log.Println("Running WDTK authentication service")

	if err := microservice.RegisterService[context.AuthData](microservice.ServiceDescription[context.AuthData]{
		SetupFn: setupAuthService,
	}, []microservice.ServiceEndpoint{
		{
			Type:            microservice.ENDPOINT_TYPE_POST,
			Name:            "Login/",
			Fn:              endpoints.Login,
			EndpointContext: nil,
			AuthRequired:    false,
		},
		{
			Type:            microservice.ENDPOINT_TYPE_POST,
			Name:            "Register/",
			Fn:              endpoints.Register,
			EndpointContext: nil,
			AuthRequired:    false,
		},
		{
			Type:            microservice.ENDPOINT_TYPE_GET,
			Name:            "Me/",
			Fn:              endpoints.Me,
			EndpointContext: nil,
			AuthRequired:    true,
		},
	}); err != nil {
		log.Println(err)
		panic("Failed to register authentication service")
	}
}
