package main

import (
	"fmt"

	"github.com/julienroland/copro/prompt"
)

func main() {
	ask := prompt.NewConfirm()
	ask.Question = "Are you a gopher"
	ask.Default = "yes"
	result := ask.Run()
	fmt.Println(result)
}
