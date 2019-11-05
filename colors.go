package tcui

import (
	"github.com/gdamore/tcell"
)

//
var (
	White      tcell.Color
	Gray       tcell.Color
	LightBlack tcell.Color
	Black      tcell.Color
	Red        tcell.Color
	Green      tcell.Color
	Blue       tcell.Color
	Yellow     tcell.Color
)

func init() {
	White = tcell.NewRGBColor(232, 234, 237)
	Gray = tcell.NewRGBColor(154, 160, 166)
	LightBlack = tcell.NewRGBColor(48, 49, 52)
	Black = tcell.NewRGBColor(32, 33, 36)
	Red = tcell.NewRGBColor(224, 108, 117)
	Green = tcell.NewRGBColor(129, 201, 149)
	Blue = tcell.NewRGBColor(138, 180, 248)
	Yellow = tcell.NewRGBColor(253, 214, 99)

}
