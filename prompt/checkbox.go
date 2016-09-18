package prompt

import (
	"fmt"

	"github.com/julienroland/usg"
)

type Checkbox struct {
	Question      string
	Choices       []*Choice
	CheckedChoice []*Choice
	app           *App
}

type CheckboxResult struct {
	ID    []int
	Label []string
}

func NewCheckbox() *Checkbox {
	CheckboxList := new(Checkbox)
	app := NewApp()
	CheckboxList.app = app
	CheckboxList.app.Run()
	return CheckboxList
}

func (s *Checkbox) Run() (CheckboxResult, error) {
	s.app.entryCount = len(s.Choices) - 1
	s.setChoicePointer()
	s.manageDefaultPointer()

	s.app.Renderer(func() {
		s.RenderHeader()
		s.RenderChoices(s.app.pointer)
	})

	ids := []int{}
	labels := []string{}
	for _, pointer := range s.app.savedPointers {
		for _, choice := range s.Choices {
			if choice.pointer == pointer {
				ids = append(ids, choice.ID)
				labels = append(labels, choice.Label)
			}
		}
	}
	result := CheckboxResult{ID: ids, Label: labels}
	return result, nil
}

func (s *Checkbox) manageDefaultPointer() {
	ids := []int{}
	for _, choice := range s.CheckedChoice {
		ids = append(ids, choice.ID)
	}
	pointers := s.searchIdInChoice(ids, s.Choices)
	s.app.savedPointers = pointers

}
func (s *Checkbox) searchIdInChoice(ids []int, choices []*Choice) []int {
	pointers := []int{}
	for _, choice := range choices {
		for _, id := range ids {
			if choice.ID == id {
				pointers = append(pointers, choice.pointer)
			}
		}
	}
	return pointers
}

func (s *Checkbox) setChoicePointer() error {
	for index, _ := range s.Choices {
		s.Choices[index].pointer = index
	}
	return nil
}

func (s *Checkbox) RenderHeader() {
	content := ""

	if len(s.Question) <= 0 {
		fmt.Errorf("You need to ask something")
	}

	content += s.Question + " \n"
	content += fmt.Sprintf("Press <%s> key to select an item", s.app.KeyboardConfig.SelectKey[0])

	display(content)
}

func (s *Checkbox) RenderChoices(currentPosition int) {

	for _, choice := range s.Choices {
		isSelected := false
		if choice.pointer == currentPosition {
			isSelected = true
		}
		isChecked := false
		for _, pointer := range s.app.savedPointers {
			if choice.pointer == pointer {
				isChecked = true
			}
		}
		line := ""
		if isSelected {
			line += usg.Get.Pointer + " "
		} else {
			line += "  "
		}

		if isChecked {
			line += usg.Get.CircleFilled + " "
		} else {
			line += usg.Get.Circle
		}

		line += choice.Label

		if isSelected {
			displayCyan(line)
			continue
		}
		if isChecked {
			displayYellow(line)
			continue
		}
		display(line)
	}

}
