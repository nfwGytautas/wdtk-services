package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nfwGytautas/wdtk-go-backend/microservice"
	"github.com/nfwGytautas/wdtk-services/auth/context"
	"golang.org/x/crypto/bcrypt"
)

type registerIn struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

func Register(executor *microservice.EndpointExecutor) {
	log.Println("Executing register request")

	var requestData registerIn
	err := json.Unmarshal(executor.Body, &requestData)
	if err != nil {
		executor.Return(http.StatusBadRequest, err)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(requestData.Password), bcrypt.MinCost)
	if err != nil {
		executor.Return(http.StatusBadRequest, err)
		return
	}

	u := context.User{
		Identifier: requestData.Identifier,
		Password:   string(hash),
		Role:       "new",
	}

	err = executor.ServiceContext.(*context.AuthData).CreateUser(&u)
	if err != nil {
		executor.Return(http.StatusInternalServerError, err)
		return
	}

	executor.Return(http.StatusNoContent, nil)
}
