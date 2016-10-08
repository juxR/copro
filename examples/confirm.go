package main

import (
	"fmt"

	"github.com/julienroland/copro/prompt"
)

func main() {
	ask := prompt.NewConfirm()
	ask.Question = "Are you a gopher"
	result := ask.Run()
	fmt.Println(result)
}
