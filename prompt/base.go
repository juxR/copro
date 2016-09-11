package prompt

import (
	"os"

	keyboard "github.com/jteeuwen/keyboard/termbox"
	term "github.com/nsf/termbox-go"
)

type App struct {
	pointer        int
	done           bool
	entryCount     int
	KeyboardConfig KeyboardConfig
	Keyboard       keyboard.Keyboard
}

type KeyboardConfig struct {
	ValidateKey       []string
	UpNavigationKey   []string
	DownNavigationKey []string
	Cancelkey         []string
}

func NewApp() *App {
	app := new(App)
	app.pointer = 0
	app.done = false
	app.KeyboardConfig = KeyboardConfig{
		ValidateKey:       []string{"enter", "space"},
		UpNavigationKey:   []string{"up", "k"},
		DownNavigationKey: []string{"down", "j"},
		Cancelkey:         []string{"ctrl+c", "esc"},
	}
	app.Keyboard = keyboard.New()

	return app
}

func (app *App) Run() {
	err := term.Init()
	if err != nil {
		panic(err)
	}
	term.HideTCursor()

	app.registerEvents()
}

func (app *App) registerEvents() {
	app.Keyboard.Bind(func() {
		os.Exit(1)
	}, app.KeyboardConfig.Cancelkey...)

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
