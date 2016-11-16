package main

var MAP_OFFSET = Size{2, 5}

type Game struct {
	Map  Map
	Turn int
}

type Map struct {
	Name string
	Pos  Point
	Size Size

	Ships        []Ship
	DisplayShips bool

	Hits   []Point
	Misses []Point

	Cursor        Point
	DisplayCursor bool
}

func NewMap(name string, pos Point, size Size) *Map {
	return &Map{
		Name: name,
		Pos:  pos,
		Size: Size{10, 10},

		Ships:        make([]Ship, 0),
		DisplayShips: false,

		Hits:   make([]Point, 0),
		Misses: make([]Point, 0),

		Cursor:        Point{0, 0},
		DisplayCursor: false,
	}
}

func (m *Map) Draw(style Style) {
	// title
	title := NewHorizontalLine(m.Pos, m.Pos.AddX(m.Size.Width))
	title.DrawCenteredString(m.Name, style)

	// grid borders
	rect := Rect{
		m.Pos.AddY(3),
		m.Size.Offset(Size{2, 2}),
	}
	rect.DrawBox(style)

	// x labels
	Line{
		Pos:      m.Pos.Offset(Point{1, 2}),
		Length:   m.Size.Width,
		Vertical: false,
	}.DrawString("0123456789", style)

	// y labels
	Line{
		Pos:      m.Pos.Offset(Point{-1, 4}),
		Length:   m.Size.Height,
		Vertical: true,
	}.DrawString("ABCDEFGHIJ", style)

	// draw ships
	if m.DisplayShips {
		for _, ship := range m.Ships {
			m.drawShip(ship, style)
		}
	}

	// draw hits
	for _, point := range m.Hits {
		point.Offset(m.Pos).AddSize(MAP_OFFSET).Draw('■', Style{Red, Default})
	}

	// draw misses
	for _, point := range m.Misses {
		point.Offset(m.Pos).AddSize(MAP_OFFSET).Draw('□', Style{White, Default})
	}

	// cursor
	if m.DisplayCursor {
		cursorPos := m.Cursor.Offset(m.Pos).AddSize(MAP_OFFSET)

		if m.overHit() {
			cursorPos.Draw('■', Style{Red, Blue})
		} else if m.overMiss() {
			cursorPos.Draw('□', Style{White, Blue})
		} else {
			cursorPos.Draw(' ', Style{Default, Blue})
		}
	}
}

func (m *Map) drawShip(ship Ship, style Style) {
	for _, point := range ship.Points {
		point.
			Offset(m.Pos).
			AddSize(MAP_OFFSET).
			Draw('■', Style{Blue, Default})
	}
}

func (m *Map) Fire() {
	if m.overShip() {
		m.hit()
	} else {
		m.miss()
	}
}

func (m *Map) overShip() bool {
	for _, ship := range m.Ships {
		for _, point := range ship.Points {
			if point.Equals(m.Cursor) {
				return true
			}
		}
	}
	return false
}

func (m *Map) overHit() bool {
	for _, point := range m.Hits {
		if point.Equals(m.Cursor) {
			return true
		}
	}
	return false
}

func (m *Map) overMiss() bool {
	for _, point := range m.Misses {
		if point.Equals(m.Cursor) {
			return true
		}
	}
	return false
}

func (m *Map) hit() {
	m.Hits = append(m.Hits, m.Cursor)
}

func (m *Map) miss() {
	m.Misses = append(m.Misses, m.Cursor)
}

type Ship struct {
	Points []Point
}

func NewShip(p Point, length int, vertical bool) Ship {
	ship := Ship{}
	ship.Points = make([]Point, length)

	for i := 0; i < length; i++ {
		if vertical {
			ship.Points[i] = p.AddY(i)
		} else {
			ship.Points[i] = p.AddX(i)
		}
	}

	return ship
}
