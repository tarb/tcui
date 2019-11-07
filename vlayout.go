package tcui

import (
	"github.com/gdamore/tcell"
)

//
type VLayout struct {
	Screen   tcell.Screen
	Children []Element

	MinWidth  int
	MinHeight int
	Border    Border
	Padding   Padding
}

//
func (vl *VLayout) Draw(x, y int, focused Element) {
	// x and y to start drawing children
	var eX, eY int = x + vl.Padding.Left() + vl.Border.Adjust(Left), y + vl.Padding.Up() + vl.Border.Adjust(Up)

	for _, e := range vl.Children {
		e.Draw(eX, eY, focused)
		var _, eHeight int = e.Size()
		eY += eHeight
	}
	vl.drawBorder(x, y)
}

//
func (vl *VLayout) Size() (int, int) {
	var cumulativeY, maxX int

	for _, e := range vl.Children {
		var w, h int = e.Size()

		cumulativeY += h
		if w > maxX {
			maxX = w
		}
	}

	if maxX < vl.MinWidth {
		maxX = vl.MinWidth
	}
	if cumulativeY < vl.MinHeight {
		cumulativeY = vl.MinHeight
	}

	cumulativeY += vl.Padding.Up() + vl.Padding.Down() + vl.Border.Adjust(Up) + vl.Border.Adjust(Down)
	maxX += vl.Padding.Left() + vl.Padding.Right() + vl.Border.Adjust(Left) + vl.Border.Adjust(Right)

	return maxX, cumulativeY
}

//
func (vl *VLayout) SetTheme(theme Theme) {
	for _, e := range vl.Children {
		e.SetTheme(theme)
	}
}

//
// func (vl *VLayout) HandleClick(mouseX, mouseY int) {
// 	fmt.Println("vlayout", vl.Border, "|", mouseX, mouseY, vl.Padding)
// }

func (vl *VLayout) drawBorder(x, y int) {
	var runes []rune = vl.Border.Runes()
	style := tcell.StyleDefault.Foreground(vl.Border.Fg).Background(vl.Border.Bg)
	w, h := vl.Size()

	// x
	if vl.Border.Has(Up) {
		for i := x; i < x+w; i++ {
			vl.Screen.SetContent(i, y, runes[0], nil, style)
		}
	}
	if vl.Border.Has(Down) {
		for i := x; i < x+w; i++ {
			vl.Screen.SetContent(i, y+h-1, runes[0], nil, style)

		}
	}
	// y
	if vl.Border.Has(Left) {
		for i := y; i < y+h; i++ {
			vl.Screen.SetContent(x, i, runes[1], nil, style)
		}
	}
	if vl.Border.Has(Right) {
		for i := y; i < y+h; i++ {
			vl.Screen.SetContent(x+w-1, i, runes[1], nil, style)
		}
	}

	// corners
	if vl.Border.Has(Left | Up) {
		vl.Screen.SetContent(x, y, runes[2], nil, style)
	}
	if vl.Border.Has(Left | Down) {
		vl.Screen.SetContent(x, y+h-1, runes[4], nil, style)
	}
	if vl.Border.Has(Right | Up) {
		vl.Screen.SetContent(x+w-1, y, runes[3], nil, style)
	}
	if vl.Border.Has(Right | Down) {
		vl.Screen.SetContent(x+w-1, y+h-1, runes[5], nil, style)
	}
}

//
func (vl *VLayout) GetFocusable() []Focusable {
	eles := make([]Focusable, 0, 10)

	for _, child := range vl.Children {
		if cont, ok := child.(Container); ok {
			eles = append(eles, cont.GetFocusable()...)
		} else if focusable, ok := child.(Focusable); ok {
			eles = append(eles, focusable)
		}
	}

	return eles
}

//
func (vl *VLayout) NextFocusable(current Focusable) Focusable {
	eles := vl.GetFocusable()

	// if there are focusable
	if len(eles) > 0 {
		//find the next focusable
		if current != nil {
			var curIdx int

			for i, e := range eles {
				if e == current {
					curIdx = i
				}
			}

			if curIdx < len(eles)-1 {
				return eles[curIdx+1]
			}
			return eles[0]

		}
		// if nothing is currently focused, return the first
		return eles[0]
	}
	// if there are no focusable, return nil
	return nil
}

//
func (vl *VLayout) FocusClicked(ev tcell.EventMouse) Focusable {
	// var w, h int = vl.Size()

	// adjust for padding
	x, y := ev.Position()
	x, y = x-vl.Padding.Left()-vl.Border.Adjust(Left), y-vl.Padding.Up()-vl.Border.Adjust(Up)
	// w, h = w-(vl.Padding.Left()+vl.Padding.Right()), h-(vl.Padding.Up()+vl.Padding.Down())

	var sumY int
	for _, c := range vl.Children {
		var cw, ch int = c.Size()

		if x >= 0 && y >= sumY && x < cw && y < sumY+ch {

			// update coords
			y -= sumY

			newEv := *tcell.NewEventMouse(x, y, tcell.Button1, 0)

			if clickable, ok := c.(Clickable); ok {
				clickable.HandleClick(newEv)
			}
			if cont, ok := c.(Container); ok {
				return cont.FocusClicked(newEv)
			} else if foc, ok := c.(Focusable); ok {
				return foc
			}
		}

		sumY += ch
	}

	return nil
}
