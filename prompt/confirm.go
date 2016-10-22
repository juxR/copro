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
	Default  string
}

func NewConfirm() *Confirm {
	return new(Confirm)
}

func (confirm *Confirm) Run() bool {
	response := false
	for {
		if len(confirm.Default) > 0 {
			fmt.Printf("%s[?]%s %s (yes/no) %s[%s]%s: ", chalk.Yellow, chalk.Green, confirm.Question, chalk.Yellow, confirm.Default, chalk.ResetColor)
		} else {
			fmt.Printf("%s[?]%s %s (yes/no)%s: ", chalk.Yellow, chalk.Green, confirm.Question, chalk.ResetColor)
		}
		userInput, err := terminal.ReadPassword(0)
		if err != nil {
			log.Fatal(err)
		}
		userResponse := string(userInput)
		userResponse = strings.ToLower(strings.TrimSpace(userResponse))

		if len(confirm.Default) > 0 && len(userResponse) <= 0 {
			userResponse = confirm.Default
		}

		if userResponse == "y" || userResponse == "yes" {
			return true
		} else if userResponse == "n" || userResponse == "no" {
			return false
		}
		fmt.Println()
	}

	return response
}
