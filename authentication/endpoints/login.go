package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nfwGytautas/gdev/jwt"
	"github.com/nfwGytautas/wdtk-go-backend/microservice"
)

type loginIn struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type loginOut struct {
	Token string `json:"token"`
}

func Login(executor *microservice.EndpointExecutor) {
	log.Println("Executing login request")

	var requestData loginIn
	err := json.Unmarshal(executor.Body, &requestData)
	if err != nil {
		executor.Return(http.StatusBadRequest, err)
		return
	}

	// TODO: Database query

	// Generate jwt token
	token, err := jwt.GenerateToken(123, "Role")
	if err != nil {
		executor.Return(http.StatusInternalServerError, err)
		return
	}

	result := loginOut{
		Token: token,
	}

	executor.Return(http.StatusOK, result)
}
