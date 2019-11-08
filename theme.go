package tcui

import "github.com/gdamore/tcell"

//
type Theme interface {
	TextCol() tcell.Color
	FocusTextCol() tcell.Color
	BoldTextCol() tcell.Color
	FocusBoldTextCol() tcell.Color
	ElementCol() tcell.Color
	FocusElementCol() tcell.Color
	BackgroundCol() tcell.Color
	CheckCol() tcell.Color
	FocusCheckCol() tcell.Color
	CursorCol() tcell.Color
	CursorBackgroundCol() tcell.Color
	TextPlaceholderCol() tcell.Color
	LoadingCol() tcell.Color
}

//
type TcuiTheme struct {
	textCol             tcell.Color
	focusTextCol        tcell.Color
	boldTextCol         tcell.Color
	focusBoldTextCol    tcell.Color
	elementCol          tcell.Color
	focusElementCol     tcell.Color
	backgroundCol       tcell.Color
	checkCol            tcell.Color
	focusCheckCol       tcell.Color
	cursorCol           tcell.Color
	cursorBackgroundCol tcell.Color
	textPlaceholderCol  tcell.Color
	loadingCol          tcell.Color
}

//
func (t TcuiTheme) TextCol() tcell.Color { return t.textCol }

//
func (t TcuiTheme) FocusTextCol() tcell.Color { return t.focusTextCol }

//
func (t TcuiTheme) BoldTextCol() tcell.Color { return t.boldTextCol }

//
func (t TcuiTheme) FocusBoldTextCol() tcell.Color { return t.focusBoldTextCol }

//
func (t TcuiTheme) ElementCol() tcell.Color { return t.elementCol }

//
func (t TcuiTheme) FocusElementCol() tcell.Color { return t.focusElementCol }

//
func (t TcuiTheme) BackgroundCol() tcell.Color { return t.backgroundCol }

//
func (t TcuiTheme) CheckCol() tcell.Color { return t.checkCol }

//
func (t TcuiTheme) FocusCheckCol() tcell.Color { return t.focusCheckCol }

//
func (t TcuiTheme) CursorCol() tcell.Color { return t.cursorCol }

//
func (t TcuiTheme) CursorBackgroundCol() tcell.Color { return t.cursorBackgroundCol }

//
func (t TcuiTheme) TextPlaceholderCol() tcell.Color { return t.textPlaceholderCol }

//
func (t TcuiTheme) LoadingCol() tcell.Color { return t.loadingCol }

//
var DefaultTheme Theme = TcuiTheme{
	textCol:             tcell.NewRGBColor(154, 160, 166),
	focusTextCol:        tcell.NewRGBColor(232, 234, 237),
	boldTextCol:         tcell.NewRGBColor(232, 234, 237),
	focusBoldTextCol:    tcell.NewRGBColor(232, 234, 237),
	elementCol:          tcell.NewRGBColor(48, 49, 52),
	focusElementCol:     tcell.NewRGBColor(48, 49, 52),
	backgroundCol:       tcell.NewRGBColor(32, 33, 36),
	checkCol:            tcell.NewRGBColor(154, 160, 166),
	focusCheckCol:       tcell.NewRGBColor(232, 234, 237),
	cursorCol:           tcell.NewRGBColor(154, 160, 166),
	cursorBackgroundCol: tcell.NewRGBColor(138, 180, 248),
	textPlaceholderCol:  tcell.NewRGBColor(154, 160, 166),
	loadingCol:          tcell.NewRGBColor(232, 234, 237),
}
