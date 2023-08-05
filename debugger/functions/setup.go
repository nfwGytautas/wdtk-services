package functions

import (
	"bufio"
	"os"

	"github.com/nfwGytautas/wdtk-go-backend/microservice"
)

var config *microservice.MicroserviceConfig
var Reader *bufio.Reader

func Setup() {
	var err error

	config, err = microservice.ReadConfig()
	if err != nil {
		panic(err)
	}

	Reader = bufio.NewReader(os.Stdin)
}
