package day9

import "math"

type Rectangle struct {
	start Position
	end   Position
}

func (r *Rectangle) Center() Position {
	return Position{y: (r.end.y + r.start.y) / 2, x: (r.start.x + r.end.x) / 2}
}

func (r *Rectangle) Area() int {
	x := math.Abs(r.start.x-r.end.x) + 1
	y := math.Abs(r.start.y-r.end.y) + 1
	return int(x * y)
}

func (r *Rectangle) GetRectByCorners() []Position {
	minX := math.Min(float64(r.start.x), float64(r.end.x))
	maxX := math.Max(float64(r.start.x), float64(r.end.x))
	minY := math.Min(float64(r.start.y), float64(r.end.y))
	maxY := math.Max(float64(r.start.y), float64(r.end.y))

	return []Position{
		{x: (minX), y: (minY)},
		{x: (minX), y: (maxY)},
		{x: (maxX), y: (minY)},
		{x: (maxX), y: (maxY)},
	}
}

