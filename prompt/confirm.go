package prompt

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ttacon/chalk"
)

type Confirm struct {
	Question string
}

func NewConfirm() *Confirm {
	return new(Confirm)
}

func (s *Confirm) Run() bool {
	var response bool
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s[?]%s %s: (Y/n) ", chalk.Yellow, chalk.ResetColor, s.Question)
		userResponse, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		userResponse = strings.ToLower(strings.TrimSpace(userResponse))

		if userResponse == "y" || userResponse == "yes" {
			return true
		} else if userResponse == "n" || userResponse == "no" {
			return false
		}
	}

	return response
}
