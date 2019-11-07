package tcui

import "github.com/gdamore/tcell"

//
type Theme struct {
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
}

//
var DefaultTheme = &Theme{
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
}
