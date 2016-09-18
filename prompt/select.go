package prompt

import (
	"fmt"

	"github.com/julienroland/usg"
)

type Select struct {
	Question       string
	Choices        []*Choice
	SelectedChoice Choice
	app            *App
}

type SelectResult struct {
	ID    int
	Label string
}

func NewSelect() *Select {
	selectList := new(Select)
	app := NewApp()
	selectList.app = app
	selectList.app.Run()
	return selectList
}

func (s *Select) Run() (SelectResult, error) {
	s.app.entryCount = len(s.Choices) - 1
	s.setChoicePointer()
	s.manageDefaultPointer()

	s.app.Renderer(func() {
		s.RenderHeader()
		s.RenderChoices(s.app.pointer)
	})

	choice := s.Choices[s.app.pointer]
	result := SelectResult{ID: choice.ID, Label: choice.Label}
	return result, nil
}

func (s *Select) manageDefaultPointer() {
	pointer, err := s.searchIdInChoice(s.SelectedChoice.ID, s.Choices)
	if err == nil {
		s.app.pointer = pointer
	}

}
func (s *Select) searchIdInChoice(id int, choices []*Choice) (int, error) {
	for _, choice := range choices {
		if choice.ID == id {
			return choice.pointer, nil
		}
	}
	return 0, fmt.Errorf("Unable to find default value: %d", id)
}

func (s *Select) setChoicePointer() error {
	for index, _ := range s.Choices {
		s.Choices[index].pointer = index
	}
	return nil
}

func (s *Select) RenderHeader() {
	content := ""

	if len(s.Question) <= 0 {
		fmt.Errorf("You need to ask something")
	}

	content += s.Question + " \n"
	content += fmt.Sprintf("Press <%s> key to select an item", s.app.KeyboardConfig.ValidateKey[0])

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
