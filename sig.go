package main

import (
	"fmt"
	"log"

	"github.com/julienroland/sig/prompt"
)

func main() {
	mylist := prompt.NewSelect()
	mylist.Question = "Choice something:"
	mylist.Choices = []*prompt.Choice{
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
	result, err := mylist.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.Label)

}
