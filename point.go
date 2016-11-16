package main

import "github.com/nsf/termbox-go"

type Point struct {
	X int
	Y int
}

func (p Point) Equals(o Point) bool {
	return p.X == o.X && p.Y == o.Y
}

func (p Point) AddSize(s Size) Point {
	return Point{
		X: p.X + s.Width - 1,
		Y: p.Y + s.Height - 1,
	}
}

func (p Point) Offset(o Point) Point {
	return Point{
		X: p.X + o.X,
		Y: p.Y + o.Y,
	}
}

func (p Point) AddX(x int) Point {
	return Point{p.X + x, p.Y}
}

func (p Point) AddY(y int) Point {
	return Point{p.X, p.Y + y}
}

func (p Point) Clamp(x, y int) Point {
	return Point{
		X: clamp(p.X, 0, x),
		Y: clamp(p.Y, 0, y),
	}
}

func (p Point) Draw(ch rune, style Style) {
	termbox.SetCell(p.X, p.Y, ch, style.Fg, style.Bg)
}
