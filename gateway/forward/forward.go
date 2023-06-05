package forward

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/nfwGytautas/gdev/jwt"
)

// PUBLIC TYPES
// ========================================================================

// PRIVATE TYPES
// ========================================================================
var locatorTable LocatorTable

// PUBLIC FUNCTIONS
// ========================================================================

/*
Setup API gateway forwarding routes
*/
func SetupRoutes(r *gin.Engine) {
	var err error
	locatorTable, err = LoadLocatorTable()
	if err != nil {
		log.Panicln("Failed to load locator table")
	}

	// Services
	gs := r.Group("/services")

	// Protect by authentication, authorization needs to be done inside handleRequest
	gs.Use(jwt.AuthenticationMiddleware())
	gs.Any("/:service/:endpoint/*params", handleServicesRequest)
}

// PRIVATE FUNCTIONS
// ========================================================================

/*
Handles a /services/:service/* request
*/
func handleServicesRequest(c *gin.Context) {
	// Get service name and endpoint
	serviceName := c.Param("service")
	endpointName := c.Param("endpoint")

	if serviceName == "" || endpointName == "" {
		c.String(http.StatusBadRequest, "service not specified")
		return
	}

	// Check in with the service locator that the service name is valid
	reroutedIp := locatorTable.Map(serviceName)
	if reroutedIp == "" {
		// Service doesn't exist
		c.Status(http.StatusBadRequest)
		return
	}

	// Endpoint and service valid proxy the request
	url, err := url.Parse(fmt.Sprintf("http://%s/%s", reroutedIp, endpointName))
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "failed to create url for proxy")
		return
	}

	log.Printf("Forwarding '%s' -> '%s'", c.Request.URL.String(), url.String())

	// Create proxy and serve it
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.Director = func(req *http.Request) {
		req.Method = c.Request.Method
		req.Header = c.Request.Header
		req.Host = url.Host
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.URL.Path = url.Path
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
