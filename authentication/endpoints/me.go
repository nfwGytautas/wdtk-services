package endpoints

import (
	"log"
	"net/http"

	"github.com/nfwGytautas/wdtk-go-backend/microservice"
	"github.com/nfwGytautas/wdtk-services/auth/context"
)

type meOut struct {
	Identifier string `json:"identifier"`
	Role       string `json:"role"`
}

func Me(executor *microservice.EndpointExecutor) {
	log.Println("Executing me request")

	user, err := executor.ServiceContext.(*context.AuthData).GetUserByID(executor.RequesterInfo.ID)
	if err != nil {
		executor.Return(http.StatusBadRequest, err)
		return
	}

	result := meOut{
		Identifier: user.Identifier,
		Role:       user.Role,
	}

	executor.Return(http.StatusOK, result)
}
