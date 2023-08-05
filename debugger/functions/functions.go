package functions

import (
	"fmt"
	"log"
	"strings"
)

// Debug function:
//   - Input string
//   - Help mode
//
// ----
//   - True if processed, false otherwise
type DebugFn func(string, bool) FnResult

type FnResult int

const (
	HANDLED   = 0
	UNHANDLED = 1
	ERROR     = 2
)

var Functions = [...]DebugFn{
	PasswordHash,
}

// TODO: Autocomplete
func ReadInput() string {
	fmt.Print("-> ")
	input, err := Reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return ""
	}

	input = strings.Replace(input, "\n", "", -1)
	return input
}
