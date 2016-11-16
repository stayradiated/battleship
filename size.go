package main

type Size struct {
	Width  int
	Height int
}

func (s Size) AddWidth(width int) Size {
	return Size{s.Width + width, s.Height}
}

func (s Size) AddHeight(height int) Size {
	return Size{s.Width, s.Height + height}
}

func (s Size) Offset(o Size) Size {
	return Size{
		Width:  s.Width + o.Width,
		Height: s.Height + o.Height,
	}
}
