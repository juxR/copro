package copro

import term "github.com/nsf/termbox-go"

type App struct {
	Pointer        int
	SavedPointers  []int
	done           bool
	EntryCount     int
	hideCursor     bool
	KeyboardConfig KeyboardConfig
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
		SelectKey:         []string{"space"},
		UpNavigationKey:   []string{"up", "k"},
		DownNavigationKey: []string{"down", "j"},
		Cancelkey:         []string{"ctrl+c", "esc"},
	}

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

}
