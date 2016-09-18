package prompt

import (
	"os"

	keyboard "github.com/jteeuwen/keyboard"
	termbox "github.com/julienroland/keyboard-termbox"
	term "github.com/nsf/termbox-go"
)

type App struct {
	pointer        int
	savedPointers  []int
	done           bool
	entryCount     int
	KeyboardConfig KeyboardConfig
	Keyboard       keyboard.Keyboard
}

type Choice struct {
	ID      int
	Label   string
	Type    string
	pointer int
}

type KeyboardConfig struct {
	ValidateKey       []string
	SelectKey         []string
	UpNavigationKey   []string
	DownNavigationKey []string
	Cancelkey         []string
}

func NewApp() *App {
	app := new(App)
	app.pointer = 0
	app.done = false
	app.KeyboardConfig = KeyboardConfig{
		ValidateKey:       []string{"enter"},
		SelectKey:         []string{"space", "o"},
		UpNavigationKey:   []string{"up", "k"},
		DownNavigationKey: []string{"down", "j"},
		Cancelkey:         []string{"ctrl+c", "esc"},
	}
	app.Keyboard = termbox.New()

	return app
}

func (app *App) Run() {
	err := term.Init()
	if err != nil {
		panic(err)
	}
	term.HideCursor()

	app.registerEvents()
}

func (app *App) registerEvents() {
	app.Keyboard.Bind(func() {
		os.Exit(1)
	}, app.KeyboardConfig.Cancelkey...)

	app.Keyboard.Bind(func() {
		exist := false
		for i, pointer := range app.savedPointers {
			if app.pointer == pointer {
				exist = true
				app.savedPointers = append(app.savedPointers[:i], app.savedPointers[i+1:]...)
			}
		}
		if !exist {
			app.savedPointers = append(app.savedPointers, app.pointer)
		}
	}, app.KeyboardConfig.SelectKey...)

	app.Keyboard.Bind(func() {
		app.done = true
	}, app.KeyboardConfig.ValidateKey...)

	app.Keyboard.Bind(func() {
		maxIndex := app.entryCount
		if app.pointer+1 > maxIndex {
			app.pointer = 0
		} else {
			app.pointer += 1
		}
	}, app.KeyboardConfig.DownNavigationKey...)

	app.Keyboard.Bind(func() {
		maxIndex := app.entryCount
		if app.pointer-1 < 0 {
			app.pointer = maxIndex
		} else {
			app.pointer -= 1
		}
	}, app.KeyboardConfig.UpNavigationKey...)
}
