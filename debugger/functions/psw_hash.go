package functions

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(input string, helpMode bool) FnResult {
	if helpMode {
		fmt.Println("<- COMMAND: psw_hash 'password'")
		return UNHANDLED
	}

	if !strings.HasPrefix(input, "psw_hash") {
		return UNHANDLED
	}

	args := strings.Split(input, " ")
	if len(args) < 2 {
		fmt.Println("<- Expected 'password'")
		return ERROR
	}

	password := args[1]

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Printf("<- %v\n", err)
		return ERROR
	}

	fmt.Printf("<- %v\n", string(hash))

	return HANDLED
}
