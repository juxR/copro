package prompt

import (
	"fmt"

	"github.com/julienroland/usg"
)

type Select struct {
	Question string
	Choices  []*Choice
}

type SelectResult struct {
	ID    int
	Label string
}

type Choice struct {
	ID      int
	Label   string
	Type    string
	pointer int
}

func NewSelect() *Select {
	selectList := new(Select)
	return selectList
}

func (s *Select) Run() (SelectResult, error) {
	app := NewApp()
	app.entryCount = len(s.Choices) - 1
	app.Run()

	app.Renderer(func() {
		s.RenderHeader(app)
		s.RenderChoices(app.pointer)
	})

	choice := s.Choices[app.pointer]
	result := SelectResult{ID: choice.ID, Label: choice.Label}
	return result, nil
}

func (s *Select) RenderHeader(app *App) {
	content := ""

	if len(s.Question) <= 0 {
		fmt.Errorf("You need to ask something")
	}

	content += s.Question + " \n"
	content += fmt.Sprintf("Press <%s> key to select an item", app.KeyboardConfig.ValidateKey[0])

	for index, _ := range s.Choices {
		s.Choices[index].pointer = index
	}

	display(content)
}

func (s *Select) RenderChoices(currentPosition int) {

	output := ""
	for _, choice := range s.Choices {
		if choice.Type == "separator" {
			output += "\n"
			display(output)
			return
		}
		isSelected := false
		if choice.pointer == currentPosition {
			isSelected = true
		}
		line := ""
		if isSelected {
			line += usg.Get.Pointer + " "
		} else {
			line += "  "
		}

		line += choice.Label

		output += line
		if isSelected {
			displayCyan(line)
			continue
		}
		display(line)
	}

}
