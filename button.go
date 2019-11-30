package tcui

import (
	"github.com/gdamore/tcell"
)

//
type Button struct {
	Screen     tcell.Screen
	Text       string
	Padding    Padding
	Submit     func()
	Theme      Theme
	DisabledFn func() bool
	Disabled   bool
	NoPad      bool
}

//
func (b *Button) Draw(x, y int, focused Element) {
	theme := b.Theme
	if theme == nil {
		theme = DefaultTheme
	}

	style1 := tcell.StyleDefault.Foreground(theme.Text()).Background(theme.Element())       // default style with text
	style2 := tcell.StyleDefault.Foreground(theme.Element()).Background(theme.Background()) // style with special width chars
	if b.isDisabled() {
		style1 = tcell.StyleDefault.Foreground(theme.DisabledText()).Background(theme.DisabledElement())
		style2 = tcell.StyleDefault.Foreground(theme.DisabledElement()).Background(theme.Background())
	} else if focused == b {
		style1 = tcell.StyleDefault.Foreground(theme.FocusText()).Background(theme.FocusElement())
		style2 = tcell.StyleDefault.Foreground(theme.FocusElement()).Background(theme.Background())
	}

	x, y = x+b.Padding.Left(), y+1 // so x ==0 && y ==0 is the location of the first char

	//draw background box
	for i := -b.Padding.Left(); i < len(b.Text)+b.Padding.Right(); i++ {

		if !b.NoPad {
			b.Screen.SetContent(x+i, y-1, '▄', nil, style2)
			b.Screen.SetContent(x+i, y+1, '▀', nil, style2)
		}

		if i >= 0 && i < len(b.Text) {
			b.Screen.SetContent(x+i, y, rune(b.Text[i]), nil, style1)
		} else {
			b.Screen.SetContent(x+i, y, ' ', nil, style1)
		}
	}

}

//
func (b *Button) Size() (int, int) {
	if b.NoPad {
		return (b.Padding.Left() + b.Padding.Right()) + len(b.Text), 1
	}
	return (b.Padding.Left() + b.Padding.Right()) + len(b.Text), 3
}

//
func (b *Button) SetTheme(theme Theme) { b.Theme = theme }

//
func (b *Button) Handle(ev tcell.EventKey) {
	if b.isDisabled() {
		return
	}

	switch ev.Key() {
	case tcell.KeyEnter:
		if b.Submit != nil {
			b.Submit()
		}
	}
}

//
func (b *Button) HandleClick(ev tcell.EventMouse) {
	if b.isDisabled() {
		return
	}

	if ev.Buttons()&tcell.Button1 != 0 {
		if x, y := ev.Position(); x >= 0 && x < b.Padding.Left()+len(b.Text)+b.Padding.Right() && y >= 0 && y < 3 {
			if b.Submit != nil {
				b.Submit()
			}
		}
	}
}

func (b *Button) isDisabled() bool { return b.Disabled || (b.DisabledFn != nil && b.DisabledFn()) }
