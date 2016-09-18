package main

import (
	"fmt"

	"github.com/julienroland/sig/prompt"
)

func main() {

	confirmPrompt := prompt.NewConfirm()
	confirmPrompt.Question = "Do you like me"

	inputPrompt := prompt.NewInput()
	inputPrompt.Required = true
	inputPrompt.Question = "What's your name"
	result := inputPrompt.Run()

	fmt.Println(result)

}
