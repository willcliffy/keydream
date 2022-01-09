package objects

type Position struct {
	X int64
	Y int64
}

func NewPosition() Position {
	// TODO - spawn in fixed location
	return Position{
		X: 0,
		Y: 0,
	}
}
