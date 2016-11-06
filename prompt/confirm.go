package prompt

import (
	"fmt"
	"log"
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
			log.Fatal(err)
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
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	setPasswordCfg := rl.GenPasswordConfig()
	setPasswordCfg.SetListener(func(line []rune, pos int, key rune) (newLine []rune, newPos int, ok bool) {
		rl.SetPrompt(question)
		rl.Refresh()
		return nil, 0, false
	})
	userResponse, err := rl.ReadPasswordWithConfig(setPasswordCfg)
	if err != nil {
		return "", err
	}
	return string(userResponse), nil
}
