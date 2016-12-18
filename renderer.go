package copro

import (
	"strings"

	runewidth "github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

var lastVerticalPosition int

func (app *App) Renderer(buffer func()) {
	for !app.done {
		lastVerticalPosition = 0
		termbox.SetOutputMode(termbox.OutputNormal)
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		buffer()
		termbox.Sync()
		app.Keyboard.Poll(termbox.PollEvent())
	}
	termbox.Close()
}

func Display(msg string) {
	printLine(msg, termbox.ColorDefault)
}

func DisplayCyan(msg string) {
	printLine(msg, termbox.ColorCyan)
}

func DisplayYellow(msg string) {
	printLine(msg, termbox.ColorYellow)
}

func DisplayBlack(msg string) {
	printLine(msg, termbox.ColorBlack)
}

func DisplayBlue(msg string) {
	printLine(msg, termbox.ColorBlue)
}

func DisplayRed(msg string) {
	printLine(msg, termbox.ColorRed)
}

func DisplayGreen(msg string) {
	printLine(msg, termbox.ColorGreen)
}

func DisplayWhite(msg string) {
	printLine(msg, termbox.ColorWhite)
}

func DisplayMajenta(msg string) {
	printLine(msg, termbox.ColorMagenta)
}

func DisplayGrey(msg string) {
	printLine(msg, termbox.ColorBlack)
}

func printLine(msg string, foreground termbox.Attribute) {
	row := strings.Split(msg, "\n")
	for _, line := range row {
		x := 0
		for _, c := range line {
			termbox.SetCell(x, lastVerticalPosition, c, foreground, termbox.ColorDefault)
			w := runewidth.RuneWidth(c)
			if w == 0 || (w == 2 && runewidth.IsAmbiguousWidth(c)) {
				w = 1
			}
			x += w
		}
		lastVerticalPosition++
	}
}
