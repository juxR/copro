package prompt

import (
	"strings"

	term "github.com/nsf/termbox-go"
)

var lastVerticalPosition int

func (app *App) Renderer(buffer func()) {
	for !app.done {
		lastVerticalPosition = 0
		term.Clear(term.ColorDefault, term.ColorDefault)
		buffer()
		term.Flush()
		app.Keyboard.Poll(term.PollEvent())
	}
	term.Close()
}

func display(msg string) {
	printLine(msg, term.ColorDefault)
}

func displayCyan(msg string) {
	printLine(msg, term.ColorCyan)
}

func displayYellow(msg string) {
	printLine(msg, term.ColorYellow)
}
func displayBlack(msg string) {
	printLine(msg, term.ColorBlack)
}

func displayBlue(msg string) {
	printLine(msg, term.ColorBlue)
}

func displayRed(msg string) {
	printLine(msg, term.ColorRed)
}

func displayGreen(msg string) {
	printLine(msg, term.ColorGreen)
}

func displayWhite(msg string) {
	printLine(msg, term.ColorWhite)
}

func printLine(msg string, foreground term.Attribute) {
	row := strings.Split(msg, "\n")
	for _, line := range row {
		x := 0
		for _, c := range line {
			term.SetCell(x, lastVerticalPosition, c, foreground, term.ColorDefault)
			x++
		}
		lastVerticalPosition++
	}
}
