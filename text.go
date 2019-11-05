package tcui

import (
	"github.com/gdamore/tcell"
)

//
type Text struct {
	Screen  tcell.Screen
	Text    string
	DText   func() string
	Width   int
	Align   Align
	Padding Padding
	BgCol   tcell.Color
	TextCol tcell.Color
}

//
func (t *Text) Draw(x, y int, focused Element) {
	style := tcell.StyleDefault.Foreground(t.TextCol).Background(t.BgCol)

	// background shading
	cw, rh := t.Size()
	for r := y; r < y+rh; r++ {
		for c := x; c < x+cw; c++ {
			t.Screen.SetContent(c, r, ' ', nil, style)
		}
	}

	if t.DText != nil {
		t.Text = t.DText()
	}

	var w int = len(t.Text)
	if t.Width != 0 && t.Width < w {
		w = t.Width
	}
	x, y = x+t.Padding.Left(), y+t.Padding.Up()

	if t.Align == AlignRight {
		x += (t.Width - w)
	} else if t.Align == AlignCenter {
		x += ((t.Width - w) / 2)
	}

	for i, c := range t.Text[:w] {
		t.Screen.SetContent(x+i, y, c, nil, style)
	}
}

//
func (t *Text) Size() (int, int) {
	var w int = len(t.Text)
	if t.Width != 0 {
		w = t.Width
	}
	return t.Padding.Left() + t.Padding.Right() + w, t.Padding.Up() + t.Padding.Down() + 1
}
