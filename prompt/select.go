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
	s.app.EntryCount = len(s.Choices) - 1
	s.setChoicePointer()
	s.manageDefaultPointer()

	s.app.Renderer(func() {
		s.RenderHeader()
		s.RenderChoices(s.app.Pointer)
	})

	choice := s.Choices[s.app.Pointer]
	result := SelectResult{ID: choice.ID, Label: choice.Label}
	return result, nil
}

func (s *Select) manageDefaultPointer() {
	for _, choice := range s.Choices {
		if choice.Selected {
			s.app.Pointer = choice.pointer
			break
		}
	}
}

func (s *Select) setChoicePointer() error {
	for index, _ := range s.Choices {
		s.Choices[index].pointer = index
	}
	return nil
}

func (s *Select) RenderHeader() {
	if len(s.Question) <= 0 {
		fmt.Errorf("You need to ask something")
	}

	copro.DisplayGreen(s.Question)
	copro.DisplayGrey(fmt.Sprintf("Press <%s> key to select an item", s.app.KeyboardConfig.ValidateKey[0]))
}

func (s *Select) RenderChoices(currentPosition int) {

	for _, choice := range s.Choices {
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
