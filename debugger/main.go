package main

import (
	"fmt"
	"log"

	"github.com/nfwGytautas/wdtk-services/debugger/functions"
)

var running bool

func main() {
	log.Println("Running WDTK debugger utilities")

	functions.Setup()

	log.Println("Ready to work! (write 'help' to get list of commands)")
	running = true

	for running {
		input := functions.ReadInput()

		if len(input) == 0 {
			continue
		}

		processInput(input)
	}

}

func processInput(input string) {
	helpMode := input == "help"

	for _, f := range functions.Functions {
		result := f(input, helpMode)

		if result != functions.UNHANDLED {
			return
		}
	}

	if !helpMode {
		fmt.Printf("<- Unknown command '%s'\n", input)
	}
}
