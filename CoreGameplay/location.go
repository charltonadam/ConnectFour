package CoreGameplay

type Direction byte

const (
	Up Direction = iota
	Down
	Right
	Left
	UpRight
	UpLeft
	DownRight
	DownLeft
)

type Location struct {
	Row, Column int
}

func NewLocation(x, y int) Location {
	return Location{
		Row:    y,
		Column: x,
	}
}

func (this Location) IsValid() bool {
	if this.Row < 0 || this.Row > NumRows-1 || this.Column < 0 || this.Column > NumColumns-1 {
		return false
	}
	return true
}

func MoveLocationInDirection(loc Location, dir Direction) Location {

	if dir == Up || dir == UpRight || dir == UpLeft {
		loc.Row++
	}
	if dir == Down || dir == DownRight || dir == DownLeft {
		loc.Row--
	}
	if dir == Right || dir == UpRight || dir == DownRight {
		loc.Column++
	}
	if dir == Left || dir == UpLeft || dir == DownLeft {
		loc.Column--
	}
	return loc
}
