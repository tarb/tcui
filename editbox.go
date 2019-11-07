package tcui

import (
	"github.com/atotto/clipboard"
	"github.com/gdamore/tcell"
)

//
type EditBox struct {
	Screen      tcell.Screen
	Width       int
	MaxSize     int
	HideContent bool
	Padding     Padding
	PlaceHolder string
	Submit      func()
	OnUpdate    func([]rune, int)
	Bind        *string
	Theme       Theme

	text      []rune
	cursorIdx int
	windowIdx int
}

//
func (eb *EditBox) Draw(x, y int, focused Element) {
	theme := eb.Theme
	if theme == nil {
		theme = DefaultTheme
	}

	x, y = x+eb.Padding.Left(), y+1 // so x ==0 && y ==0 is the location of the first char

	style1 := tcell.StyleDefault.Foreground(theme.TextCol()).Background(theme.ElementCol())       // default style with text
	style2 := tcell.StyleDefault.Foreground(theme.ElementCol()).Background(theme.BackgroundCol()) // style with special width chars
	if focused == eb {
		style1 = tcell.StyleDefault.Foreground(theme.FocusTextCol()).Background(theme.FocusElementCol())
	}
	cursorStyle := tcell.StyleDefault.Foreground(theme.CursorCol()).Background(theme.CursorBackgroundCol()) // style with special width chars

	//draw background box
	for i := -eb.Padding.Left(); i < eb.Width+eb.Padding.Right(); i++ {
		eb.Screen.SetContent(x+i, y-1, '▄', nil, style2)
		eb.Screen.SetContent(x+i, y, ' ', nil, style1)
		eb.Screen.SetContent(x+i, y+1, '▀', nil, style2)
	}

	//placeholder text
	if len(eb.text) == 0 && focused != eb {
		for i, c := range eb.PlaceHolder {
			eb.Screen.SetContent(x+i, y, c, nil, style1)
		}
	}

	var start, stop, max int

	start = eb.windowIdx

	if eb.cursorIdx == len(eb.text) {
		max = eb.Width - 1
	} else {
		max = eb.Width
	}

	if eb.windowIdx+max > len(eb.text) {
		stop = len(eb.text)
	} else {
		stop = start + max
	}

	var text []rune = eb.text[start:stop]

	for i := 0; i < len(text); i++ {
		var c rune = text[i]
		if eb.HideContent {
			c = '*'
		}

		if i+start == eb.cursorIdx && focused == eb {
			eb.Screen.SetContent(x+i, y, c, nil, cursorStyle)
		} else {
			eb.Screen.SetContent(x+i, y, c, nil, style1)
		}
	}

	if eb.cursorIdx == stop && focused == eb {
		eb.Screen.SetContent(x+stop-start, y, ' ', nil, cursorStyle)
	}
	if eb.windowIdx > 0 {
		eb.Screen.SetContent(x, y, '…', nil, style1)
	}
	if eb.windowIdx+eb.Width < len(eb.text) && (!(focused == eb) || eb.cursorIdx-eb.windowIdx != eb.Width-1) { // dont show elipise if cursor is on far right
		eb.Screen.SetContent(x+eb.Width-1, y, '…', nil, style1)
	}
}

//
func (eb *EditBox) Size() (int, int) {
	return eb.Padding.Left() + eb.Width + eb.Padding.Right(), 3
}

//
func (eb *EditBox) SetTheme(theme Theme) {
	eb.Theme = theme
}

//
func (eb *EditBox) ExpandSize() (int, int) {
	return eb.Padding.Left() + eb.Width + eb.Padding.Right(), 5
}

//
func (eb *EditBox) Handle(ev tcell.EventKey) {
	switch k := ev.Key(); k {
	case tcell.KeyEnter:
		if eb.Submit != nil {
			eb.Submit()
		}
	case tcell.KeyCtrlC:
		clipboard.WriteAll(eb.Text())
	case tcell.KeyCtrlV:
		if v, err := clipboard.ReadAll(); err == nil && v != "" {
			for _, r := range v {
				eb.insert(r)
			}
			eb.updateBind()
		}
	case tcell.KeyLeft:
		eb.cursorLeft()
	case tcell.KeyRight:
		eb.cursorRight()
	case tcell.KeyBackspace, tcell.KeyBackspace2:
		eb.backSpace()
		eb.updateBind()
	case tcell.KeyDelete, tcell.KeyCtrlD:
		eb.delete()
		eb.updateBind()
	default:
		if r := ev.Rune(); r != 0 {
			eb.insert(r)
			eb.updateBind()
		}
	}
}

//
func (eb *EditBox) HandleClick(ev tcell.EventMouse) {
	if btn := ev.Buttons(); btn&tcell.Button1 != 0 {
		// fmt.Println("editbox", mouseX, mouseY)
		x, _ := ev.Position()
		if newCIdx := eb.windowIdx + x - eb.Padding.Left(); newCIdx-eb.windowIdx == 0 {
			if eb.windowIdx > 0 {
				eb.windowIdx--
			}
			eb.cursorIdx = newCIdx
		} else if newCIdx > len(eb.text) {
			eb.cursorIdx = len(eb.text)
		} else if newCIdx-eb.windowIdx == eb.Width-1 { //&& len(eb.text) > eb.Width {
			if eb.windowIdx+eb.Width < len(eb.text) {
				eb.windowIdx++
			}
			eb.cursorIdx = newCIdx
		} else if newCIdx < len(eb.text) && newCIdx-eb.windowIdx > 0 {
			eb.cursorIdx = newCIdx
		}
	}

}

//
func (eb *EditBox) cursorLeft() {
	if eb.cursorIdx > 0 {
		eb.cursorIdx--
	}

	if eb.cursorIdx < eb.windowIdx+1 && eb.windowIdx > 0 {
		eb.windowIdx--
	}
}

//
func (eb *EditBox) cursorRight() {

	if eb.cursorIdx > eb.windowIdx+eb.Width-2 && eb.cursorIdx < len(eb.text) {
		eb.windowIdx++
	}
	if eb.cursorIdx < len(eb.text) {
		eb.cursorIdx++
	}
}

//
func (eb *EditBox) backSpace() {
	if eb.cursorIdx > 0 {
		eb.text = append(eb.text[:eb.cursorIdx-1], eb.text[eb.cursorIdx:]...)
		eb.cursorIdx--
	}

	if eb.cursorIdx-eb.windowIdx < 2 && eb.windowIdx > 0 {
		eb.windowIdx--
	}
}

//
func (eb *EditBox) delete() {
	if eb.cursorIdx < len(eb.text) {
		eb.text = append(eb.text[:eb.cursorIdx], eb.text[eb.cursorIdx+1:]...)
	}

	if eb.cursorIdx-eb.windowIdx < 2 && eb.windowIdx > 0 {
		eb.windowIdx--
	}
}

//
func (eb *EditBox) insert(r rune) {
	if len(eb.text) < eb.MaxSize || eb.MaxSize == 0 {
		if len(eb.text) == eb.cursorIdx {
			eb.text = append(eb.text, r)
			eb.cursorIdx++
		} else {
			eb.text = append(eb.text, '!')
			copy(eb.text[eb.cursorIdx+1:], eb.text[eb.cursorIdx:])
			eb.text[eb.cursorIdx] = r
			eb.cursorIdx++
		}

		if eb.OnUpdate != nil {
			eb.OnUpdate(eb.text, eb.cursorIdx-1)
		}

		if eb.cursorIdx-eb.windowIdx+1 > eb.Width {
			eb.windowIdx++
		}
	}
}

//
func (eb *EditBox) Text() string { return string(eb.text) }

//
func (eb *EditBox) updateBind() {
	if eb.Bind != nil {
		*eb.Bind = string(eb.text)
	}
}
