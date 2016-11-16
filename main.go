package main

import "github.com/nsf/termbox-go"

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)
	termbox.SetOutputMode(termbox.OutputNormal)

	ui := &Ui{}
	ui.Init()
	ui.RedrawAll(termbox.Size())

	for {
		e := termbox.PollEvent()

		if e.Type == termbox.EventResize {
			ui.RedrawAll(e.Width, e.Height)
		} else if e.Type == termbox.EventKey {

			switch e.Ch {

			// quit the game
			case 'q':
				return
			}

			switch e.Key {

			// go down one
			case termbox.KeySpace:
				ui.Fire()

			// go down one
			case termbox.KeyArrowDown:
				ui.MoveDown()

			// go left one
			case termbox.KeyArrowLeft:
				ui.MoveLeft()

				// go up one
			case termbox.KeyArrowUp:
				ui.MoveUp()

			// go right one
			case termbox.KeyArrowRight:
				ui.MoveRight()

			}

			ui.RedrawAll(termbox.Size())
		}
	}
}
