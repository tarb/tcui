package tcui

import (
	"github.com/gdamore/tcell"
)

//
type Button struct {
	Screen  tcell.Screen
	Text    string
	Padding Padding

	Submit func()
}

//
func (b *Button) Draw(x, y int, focused Element) {
	var runes = []rune(b.Text)

	style1 := tcell.StyleDefault.Foreground(Gray).Background(LightBlack)  // default style with text
	style2 := tcell.StyleDefault.Foreground(LightBlack).Background(Black) // style with special width chars
	if focused == b {
		style1 = tcell.StyleDefault.Foreground(White).Background(LightBlack)
		style2 = tcell.StyleDefault.Foreground(LightBlack).Background(Black)
	}

	x, y = x+b.Padding.Left(), y+1 // so x ==0 && y ==0 is the location of the first char

	//draw background box
	for i := -b.Padding.Left(); i < len(runes)+b.Padding.Right(); i++ {

		b.Screen.SetContent(x+i, y-1, '▄', nil, style2)
		b.Screen.SetContent(x+i, y+1, '▀', nil, style2)

		if i >= 0 && i < len(b.Text) {
			b.Screen.SetContent(x+i, y, runes[i], nil, style1)
		} else {
			b.Screen.SetContent(x+i, y, ' ', nil, style1)
		}
	}
}

//
func (b *Button) Size() (int, int) {
	return (b.Padding.Left() + b.Padding.Right()) + len(b.Text), 3
}

//
func (b *Button) Handle(ev tcell.EventKey) {
	switch ev.Key() {
	case tcell.KeyEnter:
		if b.Submit != nil {
			b.Submit()
		}
	}
}

//
func (b *Button) HandleClick(ev tcell.EventMouse) {
	//fmt.Println("button", ev.MouseX, ev.MouseY, b.Padding)
	if ev.Buttons() == tcell.Button1 {
		if x, y := ev.Position(); x >= 0 && x < b.Padding.Left()+len(b.Text)+b.Padding.Right() && y >= 0 && y < 3 {
			if b.Submit != nil {
				b.Submit()
			}
		}
	}

}
