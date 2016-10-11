package prompt

import (
	"fmt"

	"github.com/julienroland/copro"
	"github.com/julienroland/usg"
)

type Checkbox struct {
	Question      string
	Choices       []*Choice
	CheckedChoice []*Choice
	app           *copro.App
}

type CheckboxResult struct {
	ID    []int
	Label []string
}

func NewCheckbox() *Checkbox {
	CheckboxList := new(Checkbox)
	app := copro.NewApp()
	CheckboxList.app = app
	CheckboxList.app.Run()
	return CheckboxList
}

func (s *Checkbox) Run() (CheckboxResult, error) {
	s.app.EntryCount = len(s.Choices) - 1
	s.setChoicePointer()
	s.manageDefaultPointer()

	s.app.Renderer(func() {
		s.RenderHeader()
		s.RenderChoices(s.app.Pointer)
	})

	ids := []int{}
	labels := []string{}
	for _, pointer := range s.app.SavedPointers {
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
	pointers := []int{}
	for _, choice := range s.Choices {
		if choice.Selected {
			pointers = append(pointers, choice.pointer)
		}
	}
	s.app.SavedPointers = pointers

}

func (s *Checkbox) setChoicePointer() error {
	for index, _ := range s.Choices {
		s.Choices[index].pointer = index
	}
	return nil
}

func (s *Checkbox) RenderHeader() {
	if len(s.Question) <= 0 {
		fmt.Errorf("You need to ask something")
	}

	copro.DisplayGreen(s.Question)
	copro.DisplayGrey(fmt.Sprintf("Press <%s> key to select an item", s.app.KeyboardConfig.SelectKey[0]))
}

func (s *Checkbox) RenderChoices(currentPosition int) {

	for _, choice := range s.Choices {
		isSelected := false
		if choice.pointer == currentPosition {
			isSelected = true
		}
		isChecked := false
		for _, pointer := range s.app.SavedPointers {
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
			copro.DisplayYellow(line)
			continue
		}

		copro.Display(line)
	}

}
