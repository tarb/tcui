package tcui

import (
	"github.com/gdamore/tcell"
)

//
type List struct {
	Items       []Element
	BindIndex   *int
	Index       int
	WindowIndex int
	Height      int
	OnChange    func(int)
}

//
func (dl *List) Draw(x, y int, focus Element) {
	var eX, eY int = x, y
	var sumH int

	for i := dl.WindowIndex; i < len(dl.Items); i++ {
		var e = dl.Items[i]

		var listFocus Element
		if i == dl.Index {
			listFocus = e
		}

		var draw func(int, int, Element)
		var size func() (int, int)
		if ex, ok := e.(Expandable); ok && listFocus != nil {
			size = ex.ExpandSize
			draw = ex.ExpandDraw
		} else {
			size = e.Size
			draw = e.Draw
		}

		var eH int
		_, eH = size()

		if sumH+eH > dl.Height {
			break
		}

		// if container is active list item, grab first focusable element inside

		if cont, ok := e.(Container); ok {
			listFocus = cont.NextFocusable(nil)
		}

		draw(eX, eY, listFocus)
		eY, sumH = eY+eH, sumH+eH
	}
}

//
func (dl *List) Size() (int, int) {
	var maxX, sumY int

	for i := dl.WindowIndex; i < len(dl.Items); i++ {
		var e = dl.Items[i]

		var eW, eH int
		if ex, ok := e.(Expandable); ok && i == dl.Index {
			eW, eH = ex.ExpandSize()
		} else {
			eW, eH = e.Size()
		}

		if eW > maxX {
			maxX = eW
		}
		sumY += eH
		if sumY > dl.Height {
			break
		}
	}

	return maxX, dl.Height
}

//

//
func (dl *List) Handle(ev tcell.EventKey) {
	if k := ev.Key(); k == tcell.KeyUp {
		dl.scrollUp()
	} else if k == tcell.KeyDown {
		dl.scrollDown()
	} else {
		// pass event on to the next focusable thing in the item
		var e = dl.Items[dl.Index]
		if cont, ok := e.(Container); ok {
			if foc := cont.NextFocusable(nil); foc != nil {
				foc.Handle(ev)
			}
		} else if foc, ok := e.(Focusable); ok {
			foc.Handle(ev)
		}
	}
}

//
func (dl *List) HandleClick(ev tcell.EventMouse) {
	// fmt.Println("list", ev.MouseX, ev.MouseY)

	switch btn := ev.Buttons(); btn {
	case tcell.WheelDown:
		if dl.WindowIndex+dl.visibleItems() < len(dl.Items) {
			dl.WindowIndex++
		}

	case tcell.WheelUp:
		if dl.WindowIndex > 0 {
			dl.WindowIndex--
		}

	case tcell.Button1:
		var sumY int

		for i := dl.WindowIndex; i < len(dl.Items); i++ {
			var e = dl.Items[i]

			var cw, ch int
			if ex, ok := e.(Expandable); ok && i == dl.Index {
				cw, ch = ex.ExpandSize()
			} else {
				cw, ch = e.Size()
			}

			if x, y := ev.Position(); x >= 0 && y >= sumY && x < cw && y < sumY+ch {

				// adjust the event before passing it down
				y -= sumY
				newEv := *tcell.NewEventMouse(x, y, tcell.Button1, 0)

				if clickable, ok := e.(Clickable); ok {
					clickable.HandleClick(newEv)
				}
				if cont, ok := e.(Container); ok {
					cont.FocusClicked(newEv)
				}

				// fire events and update bindings
				if dl.OnChange != nil && dl.Index != i { // fire registered onChange
					dl.OnChange(i)
				}
				if dl.BindIndex != nil { // update bound index
					*dl.BindIndex = i
				}
				dl.Index = i

				// scroll the WindowIndex clicked on top|bottom element (if possible)
				if i == dl.WindowIndex && dl.WindowIndex > 0 {
					dl.WindowIndex--
				} else if dl.Index > dl.WindowIndex+dl.visibleItems()-2 && dl.WindowIndex+dl.visibleItems() < len(dl.Items) {
					dl.WindowIndex++
				}

				return
			}
			sumY += ch
		}

	}

}

//
func (dl *List) scrollDown() {
	var lastIndex = len(dl.Items) - 1

	if dl.Index < lastIndex {
		dl.Index++
		if dl.BindIndex != nil {
			*dl.BindIndex = dl.Index
		}
		if dl.OnChange != nil { // fire registered onChange
			dl.OnChange(dl.Index)
		}

		var idx = dl.Index
		if idx < lastIndex {
			idx++
		}

		var sumY int
		for i := idx; i >= 0; i-- {
			var e = dl.Items[i]

			var eh int
			if ex, ok := e.(Expandable); ok && i == dl.Index {
				_, eh = ex.ExpandSize()
			} else {
				_, eh = e.Size()
			}
			sumY += eh

			if sumY >= dl.Height {
				dl.WindowIndex = i + 1
				break
			}
			if i == dl.WindowIndex {
				break
			}
		}
	}
}

//
func (dl *List) scrollUp() {
	if dl.Index > 0 {
		dl.Index--
		if dl.BindIndex != nil {
			*dl.BindIndex = dl.Index
		}
		if dl.OnChange != nil { // fire registered onChange
			dl.OnChange(dl.Index)
		}

		if dl.Index == dl.WindowIndex && dl.Index != 0 {
			dl.WindowIndex--
		}
	}
}

//
func (dl *List) visibleItems() int {
	var sumY, count int

	for i := dl.WindowIndex; i < len(dl.Items); i++ {
		var e = dl.Items[i]

		var _, ch int
		if ex, ok := e.(Expandable); ok && i == dl.Index {
			_, ch = ex.ExpandSize()
		} else {
			_, ch = e.Size()
		}

		sumY += ch
		if sumY > dl.Height {
			break
		}
		count++
	}

	return count
}
