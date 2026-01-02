package day9

type Position struct {
	y float64
	x float64
}

func (p *Position) Minus(np *Position) Position {
	return Position{x: p.x - np.x, y: p.y - np.y}
}

