package main

import (
	"unicode/utf8"

	"github.com/nsf/termbox-go"
)

type Line struct {
	Pos      Point
	Length   int
	Vertical bool
}

func NewHorizontalLine(a, b Point) Line {
	return Line{
		Pos:      a,
		Length:   distance(a.X, b.X),
		Vertical: false,
	}
}

func NewVerticalLine(a, b Point) Line {
	return Line{
		Pos:      a,
		Length:   distance(a.Y, b.Y),
		Vertical: true,
	}
}

func (l Line) Draw(ch rune, style Style) {
	if l.Vertical {
		minY, maxY := minmax(l.Pos.Y, l.Pos.Y+l.Length)
		for y := minY; y < maxY; y++ {
			termbox.SetCell(l.Pos.X, y, ch, style.Fg, style.Bg)
		}
	} else {
		minX, maxX := minmax(l.Pos.X, l.Pos.X+l.Length)
		for x := minX; x < maxX; x++ {
			termbox.SetCell(x, l.Pos.Y, ch, style.Fg, style.Bg)
		}
	}
}

func (l Line) DrawString(text string, style Style) {

	t := []byte(text)

	if l.Vertical {
		minY, maxY := minmax(l.Pos.Y, l.Pos.Y+l.Length)
		for y := minY; y < maxY; y++ {
			var ch rune

			if len(t) > 0 {
				var size int
				ch, size = utf8.DecodeRune(t)
				t = t[size:]
			} else {
				ch = ' '
			}

			termbox.SetCell(l.Pos.X, y, ch, style.Fg, style.Bg)
		}
	} else {
		minX, maxX := minmax(l.Pos.X, l.Pos.X+l.Length)
		for x := minX; x < maxX; x++ {
			var ch rune

			if len(t) > 0 {
				var size int
				ch, size = utf8.DecodeRune(t)
				t = t[size:]
			} else {
				ch = ' '
			}

			termbox.SetCell(x, l.Pos.Y, ch, style.Fg, style.Bg)
		}
	}
}

func (l Line) DrawCenteredString(text string, style Style) {
	t := []byte(text)
	width := clamp(utf8.RuneCount(t), 0, l.Length)

	oX := (l.Length - width) / 2

	l.Draw(' ', style)

	Line{
		Point{l.Pos.X + oX, l.Pos.Y},
		width,
		false,
	}.DrawString(text, style)
}
