package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nfwGytautas/wdtk-go-backend/microservice"
	"github.com/nfwGytautas/wdtk-services/auth/context"
	"golang.org/x/crypto/bcrypt"
)

type loginIn struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type loginOut struct {
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	var requestData loginIn
	if !microservice.GinParseRequestBody(c, &requestData) {
		return
	}

	// Get user
	user, err := context.Context.GetUser(requestData.Identifier)
	if err != nil {
		// Don't send error info for security
		c.JSON(http.StatusBadRequest, microservice.EndpointError{
			Description: "Credentials invalid",
			Error:       nil,
		})
		return
	}

	// Check if correct login information
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestData.Password))
	if err != nil || err == bcrypt.ErrMismatchedHashAndPassword {
		// Don't send error info for security
		c.JSON(http.StatusBadRequest, microservice.EndpointError{
			Description: "Credentials invalid",
			Error:       nil,
		})
		return
	}

	// Generate jwt token
	token, err := microservice.GenerateToken(user.ID, user.Role)
	if err != nil {
		// Don't send error info for security
		c.JSON(http.StatusBadRequest, microservice.EndpointError{
			Description: "Failed to generate a token",
			Error:       err,
		})
		return
	}

	result := loginOut{
		Token: token,
	}
	c.JSON(http.StatusOK, result)
}
