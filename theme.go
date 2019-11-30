package tcui

import "github.com/gdamore/tcell"

//
type Theme interface {
	Text() tcell.Color
	FocusText() tcell.Color
	BoldText() tcell.Color
	FocusBoldText() tcell.Color
	Element() tcell.Color
	FocusElement() tcell.Color
	Background() tcell.Color
	Check() tcell.Color
	FocusCheck() tcell.Color
	Cursor() tcell.Color
	CursorBackground() tcell.Color
	TextPlaceholder() tcell.Color
	Loading() tcell.Color
	BorderBackground() tcell.Color
	BorderForeground() tcell.Color
	DisabledElement() tcell.Color
	DisabledText() tcell.Color
}

// TcuiTheme sample theme object implementing the Theme interface
type TcuiTheme struct {
	TextCol             tcell.Color
	FocusTextCol        tcell.Color
	BoldTextCol         tcell.Color
	FocusBoldTextCol    tcell.Color
	ElementCol          tcell.Color
	FocusElementCol     tcell.Color
	BackgroundCol       tcell.Color
	CheckCol            tcell.Color
	FocusCheckCol       tcell.Color
	CursorCol           tcell.Color
	CursorBackgroundCol tcell.Color
	TextPlaceholderCol  tcell.Color
	LoadingCol          tcell.Color
	BorderBackgroundCol tcell.Color
	BorderForegroundCol tcell.Color
	DisabledElementCol  tcell.Color
	DisabledTextCol     tcell.Color
}

//
func (t TcuiTheme) Text() tcell.Color { return t.TextCol }

//
func (t TcuiTheme) FocusText() tcell.Color { return t.FocusTextCol }

//
func (t TcuiTheme) BoldText() tcell.Color { return t.BoldTextCol }

//
func (t TcuiTheme) FocusBoldText() tcell.Color { return t.FocusBoldTextCol }

//
func (t TcuiTheme) Element() tcell.Color { return t.ElementCol }

//
func (t TcuiTheme) FocusElement() tcell.Color { return t.FocusElementCol }

//
func (t TcuiTheme) Background() tcell.Color { return t.BackgroundCol }

//
func (t TcuiTheme) Check() tcell.Color { return t.CheckCol }

//
func (t TcuiTheme) FocusCheck() tcell.Color { return t.FocusCheckCol }

//
func (t TcuiTheme) Cursor() tcell.Color { return t.CursorCol }

//
func (t TcuiTheme) CursorBackground() tcell.Color { return t.CursorBackgroundCol }

//
func (t TcuiTheme) TextPlaceholder() tcell.Color { return t.TextPlaceholderCol }

//
func (t TcuiTheme) Loading() tcell.Color { return t.LoadingCol }

//
func (t TcuiTheme) BorderBackground() tcell.Color { return t.BorderBackgroundCol }

//
func (t TcuiTheme) BorderForeground() tcell.Color { return t.BorderForegroundCol }

//
func (t TcuiTheme) DisabledElement() tcell.Color { return t.DisabledElementCol }

//
func (t TcuiTheme) DisabledText() tcell.Color { return t.DisabledTextCol }

//
var DefaultTheme Theme = TcuiTheme{
	TextCol:             tcell.NewRGBColor(154, 160, 166),
	FocusTextCol:        tcell.NewRGBColor(232, 234, 237),
	BoldTextCol:         tcell.NewRGBColor(232, 234, 237),
	FocusBoldTextCol:    tcell.NewRGBColor(232, 234, 237),
	ElementCol:          tcell.NewRGBColor(48, 49, 52),
	FocusElementCol:     tcell.NewRGBColor(48, 49, 52),
	BackgroundCol:       tcell.NewRGBColor(32, 33, 36),
	CheckCol:            tcell.NewRGBColor(154, 160, 166),
	FocusCheckCol:       tcell.NewRGBColor(232, 234, 237),
	CursorCol:           tcell.NewRGBColor(154, 160, 166),
	CursorBackgroundCol: tcell.NewRGBColor(138, 180, 248),
	TextPlaceholderCol:  tcell.NewRGBColor(154, 160, 166),
	LoadingCol:          tcell.NewRGBColor(232, 234, 237),
	BorderForegroundCol: tcell.NewRGBColor(48, 49, 52),
	BorderBackgroundCol: tcell.NewRGBColor(32, 33, 36),
	DisabledTextCol:     tcell.NewRGBColor(48, 49, 52),
	DisabledElementCol:  tcell.NewRGBColor(32, 33, 36),
}
