package main

import (
	"fmt"
	"log"

	"github.com/julienroland/sig/prompt"
)

func main() {
	selectPrompt := prompt.NewSelect()
	selectPrompt.Question = "Choice something:"
	selectPrompt.SelectedChoice = prompt.Choice{
		ID: 2,
	}
	selectPrompt.Choices = []*prompt.Choice{
		{
			ID:    1,
			Label: "Gello 1",
		},
		{
			ID:    2,
			Label: "Hello 2",
		},
		{
			ID:    3,
			Label: "Aello 3",
		},
	}
	result, err := selectPrompt.Run()
	checkboxPrompt := prompt.NewCheckbox()
	checkboxPrompt.Question = "Choice something:"
	checkboxPrompt.CheckedChoice = []*prompt.Choice{
		{
			ID: 2,
		},
		{
			ID: 4,
		},
	}
	checkboxPrompt.Choices = []*prompt.Choice{
		{
			ID:    1,
			Label: "Gello 1",
		},
		{
			ID:    2,
			Label: "Hello 2",
		},
		{
			ID:    3,
			Label: "Aello 3",
		},
		{
			ID:    4,
			Label: "Gell 4",
		},
		{
			ID:    5,
			Label: "Yeo 5",
		},
	}
	//	result, err := checkboxPrompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.Label)

}
