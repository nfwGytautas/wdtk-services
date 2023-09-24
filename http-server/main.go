package main

import (
	"log"
	"path/filepath"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/nfwGytautas/gdev/file"
	"github.com/nfwGytautas/wdtk-go-backend/microservice"
)

func main() {
	log.Println("Running WDTK http server")

	config, err := microservice.ReadConfig()
	if err != nil {
		panic(err)
	}

	htmlDir, exists := config.UserDefines["htmlDirectory"]
	if !exists {
		log.Println("Missing 'htmlDirectory' config")
		return
	}

	if !file.Exists(htmlDir.(string)) {
		log.Println("The directory '" + htmlDir.(string) + "' doesn't exist")
		return
	}

	abs, err := filepath.Abs(htmlDir.(string))
	if err != nil {
		log.Println(err)
		return
	}

	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile(abs, true)))

	httpsEnabled, httpsDefined := config.UserDefines["HTTPS"]
	if httpsDefined && httpsEnabled.(bool) {
		// HTTPS requires key and certification file
		certFile, certDefined := config.UserDefines["CertFile"]
		keyFile, keyDefined := config.UserDefines["KeyFile"]

		if !certDefined || !keyDefined {
			log.Fatal("'CertFile' and 'KeyFile' must be defined")
			return
		}

		r.RunTLS(config.RunAddress, certFile.(string), keyFile.(string))
	} else {
		// HTTP
		r.Run(config.RunAddress)
	}
}
