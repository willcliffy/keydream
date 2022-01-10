package objects

type Position struct {
	X float64
	Y float64
}

func NewPosition() Position {
	// TODO - spawn in fixed location
	return Position{
		X: 0,
		Y: 0,
	}
}
