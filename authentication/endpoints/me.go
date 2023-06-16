package endpoints

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nfwGytautas/wdtk-go-backend/microservice"
)

type meOut struct {
	Identifier string `json:"identifier"`
	Role       string `json:"role"`
}

func Me(executor *microservice.EndpointExecutor) {
	log.Println("Executing me request")

	// TODO: Database query
	fmt.Println(executor.RequesterInfo)

	result := meOut{
		Identifier: "identifier",
		Role:       executor.RequesterInfo.Role,
	}

	executor.Return(http.StatusOK, result)
}
