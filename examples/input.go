package main

import (
	"fmt"

	"github.com/julienroland/copro/prompt"
)

func main() {
	ask := prompt.NewInput()
	ask.Question = "What's your age"
	ask.Default = "18"
	ask.Validation = func(answer string) bool {
		if answer == "18" {
			return false
		}
		return true
	}
	ask.ErrorMessage = func(answer string) string {
		return "No, you don't."
	}
	ask.MaximumNumberOfTry = 2
	result, err := ask.Run()
	fmt.Println(result, err)
}
