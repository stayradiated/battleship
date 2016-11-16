package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
)

var MAP_SIZE = Size{10, 10}

type Ui struct {
	shipMap  *Map
	radarMap *Map
}

func (u *Ui) Init() {
	u.radarMap = NewMap("Radar", Point{2, 1}, MAP_SIZE)
	u.radarMap.DisplayCursor = true
	u.radarMap.Ships = []Ship{
		NewShip(Point{1, 2}, 5, false), // carrier
		NewShip(Point{4, 4}, 4, true),  // battleship
		NewShip(Point{8, 3}, 3, true),  // cruiser
		NewShip(Point{1, 0}, 3, false), // submarine
		NewShip(Point{2, 4}, 2, true),  // destroyer
	}

	u.shipMap = NewMap("Ships", Point{16, 1}, MAP_SIZE)
	u.shipMap.DisplayShips = true
	u.shipMap.Ships = []Ship{
		NewShip(Point{1, 2}, 5, false), // carrier
		NewShip(Point{4, 4}, 4, true),  // battleship
		NewShip(Point{8, 3}, 3, true),  // cruiser
		NewShip(Point{1, 0}, 3, false), // submarine
		NewShip(Point{2, 4}, 2, true),  // destroyer
	}
}

func (u *Ui) RedrawAll(w, h int) {
	termbox.Clear(Default, Default)

	style := Style{Red, Default}

	u.shipMap.Draw(style)
	u.radarMap.Draw(style)

	termbox.Flush()
}

func (u *Ui) MoveUp() {
	u.radarMap.Cursor = u.radarMap.Cursor.AddY(-1).Clamp(9, 9)
}

func (u *Ui) MoveDown() {
	u.radarMap.Cursor = u.radarMap.Cursor.AddY(1).Clamp(9, 9)
}

func (u *Ui) MoveLeft() {
	u.radarMap.Cursor = u.radarMap.Cursor.AddX(-1).Clamp(9, 9)
}

func (u *Ui) MoveRight() {
	u.radarMap.Cursor = u.radarMap.Cursor.AddX(1).Clamp(9, 9)
}

func (u *Ui) Fire() {
	u.radarMap.Fire()
	u.Opponent()
}

func (u *Ui) Opponent() {
	u.shipMap.Cursor = Point{rand.Intn(9), rand.Intn(9)}
	u.shipMap.Fire()
}
