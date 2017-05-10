package copro

import (
	"os"
	"strings"

	runewidth "github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

var lastVerticalPosition int

func (app *App) Renderer(buffer func()) {
	defer termbox.Close()

	termbox.SetOutputMode(termbox.OutputNormal)
	termbox.SetInputMode(termbox.InputEsc)
	draw(buffer)
mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {

		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			case termbox.KeyCtrlC:
				termbox.Close()
				os.Exit(1)
				break mainloop

			case termbox.KeySpace:
				app.selectCurrentPointer()

			case termbox.KeyEnter:
				break mainloop

			case termbox.KeyArrowDown:
				app.moveCurrentPointerDown()
			case termbox.KeyArrowUp:
				app.moveCurrentPointerUp()
			}
			switch ev.Ch {
			case 'k':
				app.moveCurrentPointerUp()
			case 'j':
				app.moveCurrentPointerDown()
			case 'o':
				app.selectCurrentPointer()
			}
		case termbox.EventError:
			panic(ev.Err)
		}

		draw(buffer)
	}
}

func (app *App) moveCurrentPointerUp() {

	maxIndex := app.EntryCount
	if app.Pointer-1 < 0 {
		app.Pointer = maxIndex
	} else {
		app.Pointer -= 1
	}
}

func (app *App) moveCurrentPointerDown() {
	maxIndex := app.EntryCount
	if app.Pointer+1 > maxIndex {
		app.Pointer = 0
	} else {
		app.Pointer += 1
	}
}

func (app *App) selectCurrentPointer() {
	exist := false
	for i, pointer := range app.SavedPointers {
		if app.Pointer == pointer {
			exist = true
			app.SavedPointers = append(app.SavedPointers[:i], app.SavedPointers[i+1:]...)
		}
	}
	if !exist {
		app.SavedPointers = append(app.SavedPointers, app.Pointer)
	}

}
func draw(buffer func()) {
	lastVerticalPosition = 0
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	buffer()
	termbox.Sync()
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
