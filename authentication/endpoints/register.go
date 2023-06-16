package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nfwGytautas/wdtk-go-backend/microservice"
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

	// TODO: Database query

	executor.Return(http.StatusNoContent, nil)
}
