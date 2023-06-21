package endpoints

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nfwGytautas/wdtk-go-backend/microservice"
	"github.com/nfwGytautas/wdtk-services/auth/context"
)

type meOut struct {
	Identifier string `json:"identifier"`
	Role       string `json:"role"`
}

func Me(c *gin.Context) {
	log.Println("Executing me request")

	tokenInfo, err := microservice.GetTokenInfo(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, microservice.EndpointError{
			Description: "Failed to get token info",
			Error:       err,
		})
		return
	}

	user, err := context.Context.GetUserByID(tokenInfo.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, microservice.EndpointError{
			Description: "Failed to get user",
			Error:       err,
		})
		return
	}

	result := meOut{
		Identifier: user.Identifier,
		Role:       user.Role,
	}

	c.JSON(http.StatusOK, result)
}
