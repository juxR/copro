package prompt

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/julienroland/usg"
	"github.com/ttacon/chalk"
)

type Input struct {
	Question string
	Required bool
}

func NewInput() *Input {
	input := new(Input)
	input.Required = false
	return input
}

func (s *Input) Run() string {
	response := ""
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s[?]%s %s: ", chalk.Yellow, chalk.ResetColor, s.Question)
		userResponse, err := reader.ReadString('\n')
		fmt.Println(len(userResponse))
		if s.Required && len(strings.TrimSpace(userResponse)) <= 0 {
			fmt.Printf("%s[%s]%s You can't leave this field empty%s\n", chalk.Red, usg.Get.CrossGraph, chalk.Yellow, chalk.ResetColor)
			continue
		}
		if err != nil {
			log.Fatal(err)
		}
		return userResponse
	}

	return response
}
