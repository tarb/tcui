package tcui

import (
	"time"

	"github.com/gdamore/tcell"
)

// the different frames of the loading animation
var frames = [5]string{"██  ", "▐█▌ ", " ██ ", " ▐█▌", "  ██"}

// LoadingTick - duration between repaints
var LoadingTick = 100 * time.Millisecond

//
type Loading struct {
	Screen  tcell.Screen
	Padding Padding
	Theme   Theme
}

//
func (l *Loading) Draw(x, y int, focused Element) {
	theme := l.Theme
	if theme == nil {
		theme = DefaultTheme
	}

	style := tcell.StyleDefault.Foreground(theme.Loading()).Background(theme.Background())
	x, y = x+l.Padding.Left(), y+l.Padding.Up()
	// counts 0,1,2,3,4,4,3,2,1,0 ... repeat
	n := time.Now().UnixNano() / int64(LoadingTick) % 10
	n = n + (n / 5 * (((n % 5) * -2) - 1))

	var i = 5
	for _, c := range frames[n] {
		l.Screen.SetContent(x+i, y, c, nil, style)

		i++
	}
}

//
func (l *Loading) Size() (int, int) {
	return l.Padding.Left() + 4 + l.Padding.Right(), l.Padding.Up() + 1 + l.Padding.Down()
}

//
func (l *Loading) SetTheme(theme Theme) {
	l.Theme = theme
}
