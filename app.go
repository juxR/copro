package copro

import (
	"os"

	keyboard "github.com/jteeuwen/keyboard"
	termbox "github.com/julienroland/keyboard-termbox"
	term "github.com/nsf/termbox-go"
)

type App struct {
	Pointer        int
	SavedPointers  []int
	done           bool
	EntryCount     int
	hideCursor     bool
	KeyboardConfig KeyboardConfig
	Keyboard       keyboard.Keyboard
	Width          int
	Height         int
}

type Choice struct {
	ID      int
	Label   string
	Type    string
	Pointer int
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
	app.Pointer = 0
	app.done = false
	app.hideCursor = true
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
	if app.hideCursor {
		term.HideCursor()
	}

	w, h := term.Size()
	app.Width = w
	app.Height = h
	app.registerEvents()
}

func (app *App) registerEvents() {
	app.Keyboard.Bind(func() {
		term.Close()
		os.Exit(1)
	}, app.KeyboardConfig.Cancelkey...)

	app.Keyboard.Bind(func() {
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
	}, app.KeyboardConfig.SelectKey...)

	app.Keyboard.Bind(func() {
		app.done = true
	}, app.KeyboardConfig.ValidateKey...)

	app.Keyboard.Bind(func() {
		maxIndex := app.EntryCount
		if app.Pointer+1 > maxIndex {
			app.Pointer = 0
		} else {
			app.Pointer += 1
		}
	}, app.KeyboardConfig.DownNavigationKey...)

	app.Keyboard.Bind(func() {
		maxIndex := app.EntryCount
		if app.Pointer-1 < 0 {
			app.Pointer = maxIndex
		} else {
			app.Pointer -= 1
		}
	}, app.KeyboardConfig.UpNavigationKey...)
}
