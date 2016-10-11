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
	Question           string
	Required           bool
	MaximumNumberOfTry int
}

func NewInput() *Input {
	input := new(Input)
	input.Required = false
	input.MaximumNumberOfTry = -1
	return input
}

func (input *Input) Run() (string, error) {
	response := ""
	reader := bufio.NewReader(os.Stdin)
	nbOfTry := 0
	for {
		fmt.Printf("%s[?]%s %s: ", chalk.Yellow, chalk.ResetColor, input.Question)
		userResponse, err := reader.ReadString('\n')
		userResponse = strings.Replace(strings.TrimSpace(userResponse), "\n", "", 1)
		nbOfTry++
		if input.Required && len(userResponse) <= 0 {
			if nbOfTry >= input.MaximumNumberOfTry && input.MaximumNumberOfTry != -1 {
				return "", fmt.Errorf("%s[%s]%s Number of maximun try reached %d %s\n", chalk.Red, usg.Get.CrossGraph, chalk.Yellow, nbOfTry, chalk.ResetColor)
			}
			fmt.Printf("%s[%s]%s You can't leave this field empty%s\n", chalk.Red, usg.Get.CrossGraph, chalk.Yellow, chalk.ResetColor)
			continue
		}
		if err != nil {
			log.Fatal(err)
		}
		return userResponse, nil
	}

	return response, nil
}
