package prompt

import (
	"fmt"
	"log"

	linenoise "github.com/GeertJohan/go.linenoise"
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
	nbOfTry := 0
	for {
		question := input.buildQuestion()
		userResponse, err := linenoise.Line(question)
		if err != nil {
			log.Fatal(err)
		}
		if len(input.Default) > 0 && len(userResponse) <= 0 {
			userResponse = input.Default
		}

		isValid := input.Validation(userResponse)
		if !isValid {
			if nbOfTry >= input.MaximumNumberOfTry && input.MaximumNumberOfTry != -1 {
				return "", fmt.Errorf("%s[%s]%s Number of maximun try reached %d %s\n", chalk.Red, usg.Get.CrossGraph, chalk.Yellow, nbOfTry, chalk.ResetColor)
			}
			fmt.Printf("%s[%s]%s %s %s\n", chalk.Red, usg.Get.CrossGraph, chalk.Yellow, input.ErrorMessage(userResponse), chalk.ResetColor)
			continue
		}

		nbOfTry++
		return userResponse, nil
	}

	return response, nil
}

func (input *Input) buildQuestion() string {
	hasDefaultValue := len(input.Default) > 0
	if hasDefaultValue {
		return fmt.Sprintf("%s[?]%s %s %s[%s]%s: ", chalk.Yellow, chalk.Green, input.Question, chalk.Yellow, input.Default, chalk.ResetColor)
	}
	return fmt.Sprintf("%s[?]%s %s%s: ", chalk.Yellow, chalk.Green, input.Question, chalk.ResetColor)
}
