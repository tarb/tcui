package tbui

import (
	"github.com/atotto/clipboard"
	termbox "github.com/nsf/termbox-go"
)

//
type EditBox struct {
	Width       int
	MaxSize     int
	HideContent bool
	Padding     Padding
	PlaceHolder string

	Submit   func()
	OnUpdate func([]rune, int)
	Bind     *string

	text      []rune
	cursorIdx int
	windowIdx int
}

//
func (eb *EditBox) Draw(x, y int, focused Element) {
	x, y = x+eb.Padding.Left(), y+1 // so x ==0 && y ==0 is the location of the first char

	//draw background box
	for i := -eb.Padding.Left(); i < eb.Width+eb.Padding.Right(); i++ {
		termbox.SetCell(x+i, y-1, '▄', termbox.ColorBlack|termbox.AttrBold, termbox.ColorDefault)
		termbox.SetCell(x+i, y, ' ', termbox.ColorBlack|termbox.AttrBold, termbox.ColorBlack|termbox.AttrBold)
		termbox.SetCell(x+i, y+1, '▀', termbox.ColorBlack|termbox.AttrBold, termbox.ColorDefault)
	}

	//placeholder text
	if len(eb.text) == 0 && focused != eb {
		for i, c := range eb.PlaceHolder {
			termbox.SetCell(x+i, y, c, termbox.ColorDefault, termbox.ColorBlack|termbox.AttrBold)
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
			termbox.SetCell(x+i, y, c, termbox.ColorBlack|termbox.AttrBold, termbox.ColorWhite)
		} else {
			termbox.SetCell(x+i, y, c, termbox.ColorDefault, termbox.ColorBlack|termbox.AttrBold)
		}
	}

	if eb.cursorIdx == stop && focused == eb {
		termbox.SetCell(x+stop-start, y, ' ', termbox.ColorBlack, termbox.ColorWhite)
	}
	if eb.windowIdx > 0 {
		termbox.SetCell(x, y, '…', termbox.ColorDefault, termbox.ColorBlack)
	}
	if eb.windowIdx+eb.Width < len(eb.text) && (!(focused == eb) || eb.cursorIdx-eb.windowIdx != eb.Width-1) { // dont show elipise if cursor is on far right
		termbox.SetCell(x+eb.Width-1, y, '…', termbox.ColorDefault, termbox.ColorBlack)
	}
}

//
func (eb *EditBox) Size() (int, int) {
	return eb.Padding.Left() + eb.Width + eb.Padding.Right(), 3
}

//
func (eb *EditBox) ExpandSize() (int, int) {
	return eb.Padding.Left() + eb.Width + eb.Padding.Right(), 5
}

//
func (eb *EditBox) Handle(ev termbox.Event) {
	switch ev.Key {
	case termbox.KeyEnter:
		if eb.Submit != nil {
			eb.Submit()
		}
	case termbox.KeyCtrlC:
		clipboard.WriteAll(eb.Text())
	case termbox.KeyCtrlV:
		if v, err := clipboard.ReadAll(); err == nil && v != "" {
			for _, r := range v {
				eb.insert(r)
			}
			eb.updateBind()
		}
	case termbox.KeyArrowLeft:
		eb.cursorLeft()
	case termbox.KeyArrowRight:
		eb.cursorRight()
	case termbox.KeyBackspace, termbox.KeyBackspace2:
		eb.backSpace()
		eb.updateBind()
	case termbox.KeyDelete, termbox.KeyCtrlD:
		eb.delete()
		eb.updateBind()
	case termbox.KeySpace:
		eb.insert(' ')
		eb.updateBind()
	default:
		if ev.Ch != 0 {
			eb.insert(ev.Ch)
			eb.updateBind()
		}
	}
}

//
func (eb *EditBox) HandleClick(ev termbox.Event) {

	// fmt.Println("editbox", mouseX, mouseY)
	if newCIdx := eb.windowIdx + ev.MouseX - eb.Padding.Left(); newCIdx-eb.windowIdx == 0 {
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