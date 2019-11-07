package tcui

import (
	"github.com/gdamore/tcell"
)

//
type Element interface {
	Draw(int, int, Element)
	Size() (int, int)
	SetTheme(Theme)
}

//
type Expandable interface {
	Element
	ExpandSize() (int, int)
	ExpandDraw(int, int, Element)
}

//
type Focusable interface {
	Element
	Handle(tcell.EventKey)
}

//
type Clickable interface {
	Element
	HandleClick(tcell.EventMouse)
}

//
type Container interface {
	Element
	NextFocusable(Focusable) Focusable
	GetFocusable() []Focusable
	FocusClicked(tcell.EventMouse) Focusable
}
