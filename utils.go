package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func minmax(a, b int) (min, max int) {
	if a < b {
		min, max = a, b
	} else {
		min, max = b, a
	}
	return min, max
}

func clamp(x, min, max int) int {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func distance(a, b int) int {
	return abs(b-a) + 1
}

func abs(x int) int {
	switch {
	case x < 0:
		return -x
	case x == 0:
		return 0
	}
	return x
}

func write(args ...interface{}) {
	w, h := termbox.Size()

	formatString := ""
	for range args {
		formatString += "%+v "
	}

	Line{Point{0, h - 1}, w, false}.DrawString(fmt.Sprintf(formatString, args...), Style{Default, Default})
}
