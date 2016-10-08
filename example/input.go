package main

import (
	"fmt"

	"github.com/julienroland/copro/prompt"
)

func main() {
	ask := prompt.NewInput()
	ask.Question = "What's your age"
	ask.Required = true
	ask.MaximumNumberOfTry = 2
	result, err := ask.Run()
	fmt.Println(result, err)
}
