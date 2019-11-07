package tcui

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
