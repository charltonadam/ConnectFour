package CoreGameplay

type Direction byte

const (
	up Direction = iota
	down
	right
	left
	upRight
	upLeft
	downRight
	downLeft
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

	if dir == up || dir == upRight || dir == upLeft {
		loc.Row++
	}
	if dir == down || dir == downRight || dir == downLeft {
		loc.Row--
	}
	if dir == right || dir == upRight || dir == downRight {
		loc.Column++
	}
	if dir == left || dir == upLeft || dir == downLeft {
		loc.Column--
	}
	return loc
}
