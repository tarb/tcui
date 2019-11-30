package tcui

import (
	"github.com/gdamore/tcell"
)

//
type CheckBox struct {
	Screen tcell.Screen

	Checked    bool
	Symbol     rune
	Padding    Padding
	Submit     func()
	Bind       *bool
	Theme      Theme
	Disabled   bool
	DisabledFn func() bool
}

//
func (cb *CheckBox) Draw(x, y int, focused Element) {
	theme := cb.Theme
	if theme == nil {
		theme = DefaultTheme
	}

	x, y = x+cb.Padding.Left(), y+cb.Padding.Up()

	style1 := tcell.StyleDefault.Foreground(theme.Check()).Background(theme.Element())      // default style with text
	style2 := tcell.StyleDefault.Foreground(theme.Element()).Background(theme.Background()) // style with special width chars
	if focused == cb {
		style1 = tcell.StyleDefault.Foreground(theme.FocusCheck()).Background(theme.FocusElement())
		style2 = tcell.StyleDefault.Foreground(theme.FocusElement()).Background(theme.Background())
	} else if cb.isDisabled() {
		style1 = tcell.StyleDefault.Foreground(theme.DisabledText()).Background(theme.DisabledElement())
		style2 = tcell.StyleDefault.Foreground(theme.DisabledElement()).Background(theme.Background())
	}

	mark := ' '
	if cb.Checked {
		if cb.Symbol != 0 {
			mark = cb.Symbol
		} else {
			mark = 'x'
		}
	}

	cb.Screen.SetContent(x, y, '▐', nil, style2)
	cb.Screen.SetContent(x+1, y, mark, nil, style1)
	cb.Screen.SetContent(x+2, y, '▌', nil, style2)
}

//
func (cb *CheckBox) Size() (int, int) {
	return cb.Padding.Left() + 3 + cb.Padding.Right(), cb.Padding.Up() + 1 + cb.Padding.Down()
}

//
func (cb *CheckBox) SetTheme(theme Theme) {
	cb.Theme = theme
}

//
func (cb *CheckBox) Handle(ev tcell.EventKey) {
	if ev.Rune() == ' ' {
		cb.check()
	} else if ev.Key() == tcell.KeyEnter && cb.Submit != nil {
		cb.Submit()
	}
}

//
func (cb *CheckBox) HandleClick(ev tcell.EventMouse) {
	//fmt.Println("checkbox", mouseX, mouseY, cb.Padding)
	if ev.Buttons()&tcell.Button1 != 0 {
		if x, y := ev.Position(); x >= cb.Padding.Left() && x < cb.Padding.Left()+3 && y >= cb.Padding.Up() && y < cb.Padding.Up()+1 {
			cb.check()
		}
	}
}

//
func (cb *CheckBox) check() {
	cb.Checked = !cb.Checked
	if cb.Bind != nil {
		*cb.Bind = cb.Checked
	}
}

func (cb *CheckBox) isDisabled() bool { return cb.Disabled || (cb.DisabledFn != nil && cb.DisabledFn()) }
