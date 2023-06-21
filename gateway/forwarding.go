package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/nfwGytautas/wdtk-go-backend/microservice"
)

var locatorTable LocatorTable

// Setup request forwarding
func SetupRoutes(config *microservice.MicroserviceConfig, r *gin.Engine) {
	// Prepare locator table
	locatorTable.Parse(config)

	gs := r.Group("/")

	gs.Any("/:service/:endpoint/*params", handleRequest)
}

func handleRequest(c *gin.Context) {
	// Get service name and endpoint
	serviceName := c.Param("service")
	endpointName := c.Param("endpoint")

	if serviceName == "" || endpointName == "" {
		c.JSON(http.StatusBadRequest, microservice.EndpointError{
			Description: "Service or endpoint not specified",
			Error:       nil,
		})
		return
	}

	// Get ip to proxy to
	ip, err := locatorTable.GetIp(serviceName)
	if err != nil {
		c.JSON(http.StatusBadRequest, microservice.EndpointError{
			Description: "Service doesn't exist in the locator table",
			Error:       err,
		})
		return
	}

	// Endpoint and service valid proxy the request
	url, err := url.Parse(fmt.Sprintf("http://%s/%s", ip, endpointName))
	if err != nil {
		c.JSON(http.StatusBadRequest, microservice.EndpointError{
			Description: "Failed to create url object for proxy",
			Error:       err,
		})
		return
	}

	log.Printf("Forwarding '%s' -> '%s'", c.Request.URL.String(), url.String())

	// Create proxy and serve it
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.Director = func(req *http.Request) {
		req.Method = c.Request.Method
		req.Header = c.Request.Header
		req.Body = c.Request.Body
		req.Host = url.Host
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.URL.Path = url.Path
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
