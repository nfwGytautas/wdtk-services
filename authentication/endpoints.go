package main

import (
	"log"
	"net/http"

	"github.com/nfwGytautas/wdtk-go-backend/microservice"
)

type AuthData struct {
}

func loginEndpoint(executor *microservice.EndpointExecutor) {
	log.Println("Executing login request")
	executor.Return(http.StatusOK, nil)
}
