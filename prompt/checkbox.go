package prompt

import (
	"fmt"
	"math"

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
	ID    int
	Label string
}

func NewCheckbox() *Checkbox {
	CheckboxList := new(Checkbox)
	app := copro.NewApp()
	CheckboxList.app = app
	CheckboxList.app.Run()
	return CheckboxList
}

func (s *Checkbox) Run() ([]CheckboxResult, error) {
	s.app.EntryCount = len(getChoicesWithoutSeparator(s.Choices)) - 1
	s.setChoicePointer()
	s.manageDefaultPointer()
	s.app.Renderer(func() {
		s.RenderHeader()
		s.RenderChoices(s.app.Pointer)
	})

	result := []CheckboxResult{}
	for _, pointer := range s.app.SavedPointers {
		for _, choice := range s.Choices {
			if choice.pointer == pointer && !choice.IsSeparator {
				result = append(result, CheckboxResult{
					ID:    choice.ID,
					Label: choice.Label,
				})
			}
		}
	}
	return result, nil
}

func (s *Checkbox) manageDefaultPointer() {
	pointers := []int{}
	for _, choice := range s.Choices {
		if choice.Selected && !choice.IsSeparator {
			pointers = append(pointers, choice.pointer)
		}
	}
	s.app.SavedPointers = pointers

}

func (s *Checkbox) setChoicePointer() error {
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

func (s *Checkbox) RenderHeader() {
	if len(s.Question) <= 0 {
		fmt.Errorf("You need to ask something")
	}

	copro.DisplayGreen(s.Question)
	copro.Display(fmt.Sprintf("Press <%s> key to select an item", s.app.KeyboardConfig.SelectKey[0]))
}

func (s *Checkbox) RenderChoices(currentPosition int) {

	headSize := 2
	perPage := s.app.Height - headSize - 1

	nbOfSeparator := 0
	for _, choice := range s.Choices {
		if choice.IsSeparator {
			nbOfSeparator++
		}
	}
	limit := len(s.Choices)
	pages := int(math.Ceil(float64(limit) / float64(perPage)))
	current := 0
	if currentPosition >= (perPage - nbOfSeparator) {
		current = currentPosition
	}
	currentMax := int(math.Min(float64(limit), float64(current+perPage)))
	for _, choice := range s.Choices[current:currentMax] {
		if choice.IsSeparator {
			copro.DisplayGrey(" --- " + choice.Label + " ---")
			continue
		}
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
			line += usg.Get.Circle + " "
		}

		line += " " + choice.Label

		if isSelected {
			copro.DisplayYellow(line)
			continue
		}

		copro.Display(line)
	}
	if pages > 1 && current != (limit-nbOfSeparator-1) {
		copro.DisplayGrey(" --- more available ---")
	}
}
