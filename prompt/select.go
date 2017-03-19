package prompt

import (
	"fmt"

	"github.com/julienroland/copro"
	"github.com/julienroland/usg"
)

type Select struct {
	Question       string
	Choices        []*Choice
	SelectedChoice Choice
	app            *copro.App
}

type SelectResult struct {
	ID    int
	Label string
}

func NewSelect() *Select {
	selectList := new(Select)
	app := copro.NewApp()
	selectList.app = app
	selectList.app.Run()
	return selectList
}

func (s *Select) Run() (SelectResult, error) {
	choicesWithoutSeparator := getChoicesWithoutSeparator(s.Choices)
	s.app.EntryCount = len(choicesWithoutSeparator) - 1
	s.setChoicePointer()
	s.manageDefaultPointer()

	s.app.Renderer(func() {
		s.RenderHeader()
		s.RenderChoices(s.app.Pointer)
	})

	choice := choicesWithoutSeparator[s.app.Pointer]
	result := SelectResult{ID: choice.ID, Label: choice.Label}
	return result, nil
}

func (s *Select) manageDefaultPointer() {
	for _, choice := range s.Choices {
		if choice.Selected && !choice.IsSeparator {
			s.app.Pointer = choice.pointer
			break
		}
	}
}

func (s *Select) setChoicePointer() error {
	position := 0
	for _, choice := range s.Choices {
		if choice.IsSeparator {
			continue
		}
		choice.pointer = position
		position++
	}
	return nil
}

func (s *Select) RenderHeader() {
	if len(s.Question) <= 0 {
		fmt.Errorf("You need to ask something")
	}

	copro.DisplayGreen(s.Question)
	copro.Display(fmt.Sprintf("Press <%s> key to select an item", s.app.KeyboardConfig.ValidateKey[0]))
}

func (s *Select) RenderChoices(currentPosition int) {

	for _, choice := range s.Choices {
		if choice.IsSeparator {
			copro.DisplayGrey(" --- " + choice.Label + " ---")
			continue
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

		if isSelected {
			copro.DisplayYellow(line)
			continue
		}

		copro.Display(line)
	}

}

func getChoicesWithoutSeparator(allChoices []*Choice) []*Choice {
	var choices []*Choice
	for _, choice := range allChoices {
		if choice.IsSeparator {
			continue
		}
		choices = append(choices, choice)
	}
	return choices
}
