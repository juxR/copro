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
			ID:    1,
			Label: "Apple",
		},
		{
			ID:       2,
			Label:    "Tomato",
			Selected: true,
		},
		{
			ID:    3,
			Label: "Orange",
		},
		{
			ID:       4,
			Label:    "Watermelon",
			Selected: true,
		},
	}
	result, err := ask.Run()
	fmt.Println(result, err)
}