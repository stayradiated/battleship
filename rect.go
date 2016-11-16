package main

import "unicode/utf8"

type Rect struct {
	Pos  Point
	Size Size
}

func NewRect(x, y, w, h int) Rect {
	return Rect{
		Pos:  Point{x, y},
		Size: Size{w, h},
	}
}

func (r Rect) Contains(p Point) bool {
	return p.X >= r.Pos.X &&
		p.X < r.Pos.X+r.Size.Width &&
		p.Y >= r.Pos.Y &&
		p.Y < r.Pos.Y+r.Size.Height
}

func (r Rect) Fill(ch rune, style Style) {
	a, b := r.Pos, r.Pos.AddSize(r.Size)
	minY, maxY := minmax(a.Y, b.Y)
	minX, maxX := minmax(b.X, a.X)

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			Point{x, y}.Draw(ch, style)
		}
	}
}

func (r Rect) Clear() {
	r.Fill(' ', Style{Default, Default})
}

func (r Rect) DrawBox(style Style) {
	tl := r.Pos
	br := r.Pos.AddSize(r.Size)
	tr := Point{br.X, tl.Y}
	bl := Point{tl.X, br.Y}

	NewHorizontalLine(tl, tr).Draw('─', style)
	NewHorizontalLine(bl, br).Draw('─', style)
	NewVerticalLine(tl, bl).Draw('│', style)
	NewVerticalLine(tr, br).Draw('│', style)

	tl.Draw('┌', style)
	tr.Draw('┐', style)
	bl.Draw('└', style)
	br.Draw('┘', style)
}

func (r Rect) DrawCenteredString(text string, style Style) {
	t := []byte(text)
	width := clamp(utf8.RuneCount(t), 0, r.Size.Width)

	oX := (r.Size.Width - width) / 2
	oY := r.Size.Height / 2

	Line{r.Pos.AddX(oX).AddY(oY), width, false}.DrawString(text, style)
}
