package prompt

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/ttacon/chalk"
)

type Confirm struct {
	Question string
}

func NewConfirm() *Confirm {
	return new(Confirm)
}

func (confirm *Confirm) Run() bool {
	response := false
	for {
		fmt.Printf("%s[?]%s %s: %s(Y/n)%s ", chalk.Yellow, chalk.ResetColor, confirm.Question, chalk.Red, chalk.ResetColor)
		userInput, err := terminal.ReadPassword(0)
		userResponse := string(userInput)
		if err != nil {
			log.Fatal(err)
		}

		userResponse = strings.ToLower(strings.TrimSpace(userResponse))

		if userResponse == "y" || userResponse == "yes" {
			return true
		} else if userResponse == "n" || userResponse == "no" {
			return false
		}
		fmt.Println()
	}

	return response
}
