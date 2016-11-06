package main

import (
	"fmt"

	"github.com/julienroland/copro/prompt"
)

func main() {
	ask := prompt.NewPassword()
	ask.Question = "Enter your password"
	ask.Validation = func(answer string) bool {
		if answer != "password" {
			return false
		}
		return true
	}
	ask.ErrorMessage = func(answer string) string {
		return "Wrong password"
	}
	ask.MaximumNumberOfTry = 2
	result, err := ask.Run()
	fmt.Println(result, err)
}
