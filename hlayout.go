package tcui

import (
	"github.com/gdamore/tcell"
)

//
type HLayout struct {
	Screen   tcell.Screen
	Children []Element

	MinHeight int
	MinWidth  int
	Border    Border
	Padding   Padding
}

//
func (hl *HLayout) Draw(x, y int, focused Element) {
	// x and y to start drawing children
	var eX, eY int = x + hl.Padding.Left() + hl.Border.Adjust(Left), y + hl.Padding.Up() + hl.Border.Adjust(Up)

	for _, e := range hl.Children {
		e.Draw(eX, eY, focused)
		var eWidth, _ int = e.Size()
		eX += eWidth
	}
	hl.drawBorder(x, y)
}

//
func (hl *HLayout) Size() (int, int) {
	var cumulativeX, maxY int

	for _, e := range hl.Children {
		var w, h int = e.Size()

		cumulativeX += w
		if h > maxY {
			maxY = h
		}
	}

	if cumulativeX < hl.MinWidth {
		cumulativeX = hl.MinWidth
	}
	if maxY < hl.MinHeight {
		maxY = hl.MinHeight
	}

	cumulativeX += hl.Padding.Left() + hl.Padding.Right() + hl.Border.Adjust(Left) + hl.Border.Adjust(Right)
	maxY += hl.Padding.Up() + hl.Padding.Down() + hl.Border.Adjust(Up) + hl.Border.Adjust(Down)

	return cumulativeX, maxY
}

//
func (hl *HLayout) SetTheme(theme Theme) {
	for _, e := range hl.Children {
		e.SetTheme(theme)
	}
}

//
// func (hl *HLayout) HandleClick(mouseX, mouseY int) {
// 	fmt.Println("hlayout", hl.Border, "|", mouseX, mouseY, hl.Padding)
// }

func (hl *HLayout) drawBorder(x, y int) {
	var runes []rune = hl.Border.Runes()
	style := tcell.StyleDefault.Foreground(hl.Border.Fg).Background(hl.Border.Bg)
	w, h := hl.Size()

	// x
	if hl.Border.Has(Up) {
		for i := x; i < x+w; i++ {
			hl.Screen.SetContent(i, y, runes[0], nil, style)
		}
	}
	if hl.Border.Has(Down) {
		for i := x; i < x+w; i++ {
			hl.Screen.SetContent(i, y+h-1, runes[0], nil, style)

		}
	}
	// y
	if hl.Border.Has(Left) {
		for i := y; i < y+h; i++ {
			hl.Screen.SetContent(x, i, runes[1], nil, style)
		}
	}
	if hl.Border.Has(Right) {
		for i := y; i < y+h; i++ {
			hl.Screen.SetContent(x+w-1, i, runes[1], nil, style)
		}
	}

	// corners
	if hl.Border.Has(Left | Up) {
		hl.Screen.SetContent(x, y, runes[2], nil, style)
	}
	if hl.Border.Has(Left | Down) {
		hl.Screen.SetContent(x, y+h-1, runes[4], nil, style)
	}
	if hl.Border.Has(Right | Up) {
		hl.Screen.SetContent(x+w-1, y, runes[3], nil, style)
	}
	if hl.Border.Has(Right | Down) {
		hl.Screen.SetContent(x+w-1, y+h-1, runes[5], nil, style)
	}
}

//
func (hl *HLayout) GetFocusable() []Focusable {
	var eles = make([]Focusable, 0, 10)

	for _, child := range hl.Children {
		if cont, ok := child.(Container); ok {
			eles = append(eles, cont.GetFocusable()...)
		} else if focusable, ok := child.(Focusable); ok {
			eles = append(eles, focusable)
		}
	}

	return eles
}

//
func (hl *HLayout) NextFocusable(current Focusable) Focusable {
	var eles []Focusable = hl.GetFocusable()

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
func (hl *HLayout) FocusClicked(ev tcell.EventMouse) Focusable {
	// var w, h int = hl.Size()

	// termbox uses coords based from 1, 1 not 0, 0
	// keep it consistent for bubbling through containers
	// but -1,-1 for the HandleClick methods

	// if i ever add the Clickable interface to HLayout
	// if mouseX > 0 && mouseY > 0 && mouseX <= w && mouseY <= h {
	// 	hl.HandleClick(mouseX, mouseY)
	// }

	// normalise mouse click to this element so it can be
	// passed down to children

	// adjust for padding
	x, y := ev.Position()
	x, y = x-hl.Padding.Left()-hl.Border.Adjust(Left), y-hl.Padding.Up()-hl.Border.Adjust(Up)
	// w, h = w-(hl.Padding.Left()+hl.Padding.Right())-(hl.Border.Left()+hl.Border.Right()), h-(hl.Padding.Up()+hl.Padding.Down())-(hl.Border.Up()+hl.Border.Down())

	var sumX int
	for _, c := range hl.Children {
		var cw, ch int = c.Size()

		if x >= sumX && y >= 0 && x <= sumX+cw && y <= ch {
			// update event before passing it down
			x -= sumX

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

		sumX += cw
	}

	return nil
}
