package prompt

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/julienroland/usg"
	"github.com/ttacon/chalk"
)

type Input struct {
	Question           string
	Default            string
	MaximumNumberOfTry int
	ErrorMessage       func(string) string
	Validation         func(string) bool
}

func NewInput() *Input {
	input := new(Input)
	input.Validation = func(answer string) bool {
		return true
	}
	input.ErrorMessage = func(response string) string {
		return "You can't leave this field empty"
	}
	input.MaximumNumberOfTry = -1
	return input
}

func (input *Input) Run() (string, error) {
	response := ""
	reader := bufio.NewReader(os.Stdin)
	nbOfTry := 0
	for {
		if len(input.Default) > 0 {
			fmt.Printf("%s[?]%s %s %s[%s]%s: ", chalk.Yellow, chalk.Green, input.Question, chalk.Yellow, input.Default, chalk.ResetColor)
		} else {
			fmt.Printf("%s[?]%s %s%s: ", chalk.Yellow, chalk.Green, input.Question, chalk.ResetColor)
		}
		userResponse, err := reader.ReadString('\n')
		userResponse = strings.Replace(strings.TrimSpace(userResponse), "\n", "", 1)
		if len(input.Default) > 0 && len(userResponse) <= 0 {
			userResponse = input.Default
		}
		nbOfTry++
		isValid := input.Validation(userResponse)
		if !isValid {
			if nbOfTry >= input.MaximumNumberOfTry && input.MaximumNumberOfTry != -1 {
				return "", fmt.Errorf("%s[%s]%s Number of maximun try reached %d %s\n", chalk.Red, usg.Get.CrossGraph, chalk.Yellow, nbOfTry, chalk.ResetColor)
			}
			fmt.Printf("%s[%s]%s %s %s\n", chalk.Red, usg.Get.CrossGraph, chalk.Yellow, input.ErrorMessage(userResponse), chalk.ResetColor)
			continue
		}
		if err != nil {
			log.Fatal(err)
		}
		return userResponse, nil
	}

	return response, nil
}
