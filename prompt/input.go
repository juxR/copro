package prompt

import (
	"fmt"
	"io"
	"strings"

	"github.com/chzyer/readline"
	"github.com/julienroland/usg"
	"github.com/ttacon/chalk"
)

type Input struct {
	Question           string
	Default            string
	MaximumNumberOfTry int
	ErrorMessage       func(string) string
	Validation         func(string) bool
	nbOfTry            int
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
	input.nbOfTry = 0
	return input
}

func (input *Input) Run() (string, error) {
	for {
		question := input.buildQuestion()
		userResponse, err := input.readUserInput(question)
		if err != nil {
			if err == readline.ErrInterrupt {
				return "", fmt.Errorf("Question closed by the user")
			} else if err == io.EOF {
				break
			}
			return "", err
		}
		if len(input.Default) > 0 && len(userResponse) <= 0 {
			userResponse = input.Default
		}

		isValid, err := input.validate(userResponse)
		if err != nil {
			return "", err
		}
		if !isValid {
			continue
		}
		return strings.TrimSpace(userResponse), nil
	}
	return "", nil
}

func (input *Input) validate(userResponse string) (bool, error) {
	input.nbOfTry++
	isValid := input.Validation(userResponse)
	if !isValid {
		if input.nbOfTry >= input.MaximumNumberOfTry && input.MaximumNumberOfTry != -1 {
			return true, fmt.Errorf("%s[%s]%s Number of maximun try reached %d %s\n", chalk.Red, usg.Get.CrossGraph, chalk.Yellow, input.nbOfTry, chalk.ResetColor)
		}
		fmt.Printf("%s[%s]%s %s %s\n", chalk.Red, usg.Get.CrossGraph, chalk.Yellow, input.ErrorMessage(userResponse), chalk.ResetColor)
		return false, nil
	}
	return true, nil

}
func (input *Input) readUserInput(question string) (string, error) {
	rl, err := readline.NewEx(&readline.Config{
		VimMode: false,
		Prompt:  question,
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()
	userResponse, err := rl.Readline()
	if err != nil {
		return "", err
	}
	userResponse = strings.TrimSpace(userResponse)
	return userResponse, nil
}

func (input *Input) buildQuestion() string {
	hasDefaultValue := len(input.Default) > 0
	if hasDefaultValue {
		return fmt.Sprintf("%s[?]%s %s %s[%s]%s: ", chalk.Yellow, chalk.Green, input.Question, chalk.Yellow, input.Default, chalk.ResetColor)

	}
	return fmt.Sprintf("%s[?]%s %s%s: ", chalk.Yellow, chalk.Green, input.Question, chalk.ResetColor)
}
