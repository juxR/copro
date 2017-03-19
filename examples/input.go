package main

import (
	"fmt"
	"strconv"

	"github.com/julienroland/copro/prompt"
)

func main() {
	ask := prompt.NewInput()
	ask.Question = "🔒  What's your age?"
	ask.Default = "18"
	ask.Validation = func(answer string) bool {
		_, err := strconv.Atoi(answer)
		if answer == "18" || err != nil {
			return false
		}
		return true
	}
	ask.ErrorMessage = func(answer string) string {
		_, err := strconv.Atoi(answer)
		if err != nil {
			return "Only numbers are authorized"
		}

		return "No, you don't."
	}
	ask.MaximumNumberOfTry = 2
	result, err := ask.Run()
	fmt.Println(result, err)
}
