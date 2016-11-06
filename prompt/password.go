package prompt

import (
	"fmt"

	"github.com/chzyer/readline"
)

type Password struct {
	*Input
}

func NewPassword() *Password {
	input := NewInput()
	password := &Password{input}
	return password
}

func (password *Password) Run() (string, error) {
	for {
		question := password.buildQuestion()
		userResponse, err := password.readUserInput(question)
		if err != nil {
			return "", err
		}
		if len(password.Default) > 0 && len(userResponse) <= 0 {
			userResponse = password.Default
		}

		isValid, err := password.validate(userResponse)
		if err != nil {
			return "", err
		}
		if !isValid {
			continue
		}
		return userResponse, nil
	}
	return "", nil
}

func (password *Password) readUserInput(question string) (string, error) {
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
		passwordLength := ""
		for range line {
			passwordLength += "*"
		}
		fmt.Print(passwordLength)
		return nil, 0, false
	})
	userResponse, err := rl.ReadPasswordWithConfig(setPasswordCfg)
	if err != nil {
		return "", err
	}
	return string(userResponse), nil
}
