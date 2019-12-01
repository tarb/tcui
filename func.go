package tcui

import "github.com/gdamore/tcell"

//
type Func struct {
	F     func(Theme) Element
	Theme Theme
}

//
func (f *Func) Draw(x, y int, focused Element) {
	var e = f.F(f.Theme)
	if ele, ok := e.(Element); ok {
		ele.Draw(x, y, focused)
	}
}

//
func (f *Func) Size() (int, int) {
	var e = f.F(f.Theme)
	if ele, ok := e.(Element); ok {
		return ele.Size()
	}
	return 0, 0
}

//
func (f *Func) SetTheme(theme Theme) {
	f.Theme = theme
}

//
func (f *Func) GetFocusable() []Focusable {
	var eles []Focusable

	child := f.F(f.Theme)
	if cont, ok := child.(Container); ok {
		eles = cont.GetFocusable()
	} else if focusable, ok := child.(Focusable); ok {
		eles = []Focusable{focusable}
	}

	return eles
}

//
func (f *Func) NextFocusable(current Focusable) Focusable {
	eles := f.GetFocusable()

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
func (f *Func) FocusClicked(ev tcell.EventMouse) Focusable {
	c := f.F(f.Theme)

	if clickable, ok := c.(Clickable); ok {
		clickable.HandleClick(ev)
	}

	if cont, ok := c.(Container); ok {
		return cont.FocusClicked(ev)
	} else if foc, ok := c.(Focusable); ok {
		return foc
	}

	return nil
}
