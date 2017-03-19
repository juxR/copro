package main

import (
	"fmt"

	"github.com/julienroland/copro/prompt"
)

func main() {
	ask := prompt.NewCheckbox()
	ask.Question = "[?] Choose what you want:"
	ask.Choices = []*prompt.Choice{
		{
			Label:       "List",
			IsSeparator: true,
		},
		{
			ID:    1,
			Label: "Apple",
		},
		{
			ID:       2,
			Label:    "Tomato",
			Selected: true,
		},
	}
	result, err := ask.Run()
	fmt.Println(result, err)
}
