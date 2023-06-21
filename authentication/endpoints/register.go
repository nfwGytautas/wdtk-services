package endpoints

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nfwGytautas/wdtk-go-backend/microservice"
	"github.com/nfwGytautas/wdtk-services/auth/context"
	"golang.org/x/crypto/bcrypt"
)

type registerIn struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

func Register(c *gin.Context) {
	log.Println("Executing register request")

	var requestData registerIn
	if !microservice.GinParseRequestBody(c, &requestData) {
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(requestData.Password), bcrypt.MinCost)
	if err != nil {
		// Don't send error info for security
		c.JSON(http.StatusBadRequest, microservice.EndpointError{
			Description: "Failed to create password hash",
			Error:       nil,
		})
		return
	}

	u := context.User{
		Identifier: requestData.Identifier,
		Password:   string(hash),
		Role:       "new",
	}

	err = context.Context.CreateUser(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, microservice.EndpointError{
			Description: "Failed to create user",
			Error:       err,
		})
		return
	}

	c.Status(http.StatusNoContent)
}
