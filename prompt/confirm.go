package prompt

import (
	"fmt"
	"os"
	"strings"

	"github.com/chzyer/readline"
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
		question := fmt.Sprintf("%s[?]%s %s (yes/no)%s: ", chalk.Yellow, chalk.Green, confirm.Question, chalk.ResetColor)
		if len(confirm.Default) > 0 {
			question = fmt.Sprintf("%s[?]%s %s (yes/no) %s[%s]%s: ", chalk.Yellow, chalk.Green, confirm.Question, chalk.Yellow, confirm.Default, chalk.ResetColor)
		}
		userResponse, err := readUserInput(question)
		if err != nil {
			if err == readline.ErrInterrupt {
				fmt.Fprintln(os.Stderr, "Question closed by the user")
				os.Exit(1)
			}
			fmt.Fprintln(os.Stderr, fmt.Sprint(err))
			os.Exit(1)
		}
		userResponse = strings.ToLower(strings.TrimSpace(userResponse))

		if len(confirm.Default) > 0 && len(userResponse) <= 0 {
			userResponse = confirm.Default
		}

		if userResponse == "y" || userResponse == "yes" {
			response = true
			break
		} else if userResponse == "n" || userResponse == "no" {
			response = false
			break
		}
	}

	return response
}

func readUserInput(question string) (string, error) {
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
	return userResponse, nil
}
