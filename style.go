package main

import "github.com/nsf/termbox-go"

type Style struct {
	Fg termbox.Attribute
	Bg termbox.Attribute
}

func (s Style) Inverted() Style {
	return Style{
		Fg: s.Fg | Reverse,
		Bg: s.Bg | Reverse,
	}
}

func (s Style) Brightened() Style {
	return Style{
		Fg: s.Fg | Bold,
		Bg: s.Bg,
	}
}

var (
	Default = termbox.ColorDefault
	Black   = termbox.ColorBlack
	Red     = termbox.ColorRed
	Green   = termbox.ColorGreen
	Yellow  = termbox.ColorYellow
	Blue    = termbox.ColorBlue
	Magenta = termbox.ColorMagenta
	Cyan    = termbox.ColorCyan
	White   = termbox.ColorWhite

	Bold    = termbox.AttrBold
	Reverse = termbox.AttrReverse
)
