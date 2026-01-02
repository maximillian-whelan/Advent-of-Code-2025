package day9

type BoundingBox struct {
	corners []Position
}

func (bb *BoundingBox) RectInsideBoundingBox(r Rectangle) bool {
	corners := r.GetRectByCorners()
	return bb.AllCornersInside(corners) && !bb.EdgeIntersects(corners)
}

// Raycasting algorithm to define if a point is inside bounding box
// focuses on polygon shapes as likely the shape given by inputs for AoC
func (bb *BoundingBox) IsPointInside(point Position) bool {
	inside := false
	for i, pos := range bb.corners {
		j := (i + 1) % len(bb.corners) // wrap back around for final check
		xi, yi := pos.x, pos.y
		xj, yj := bb.corners[j].x, bb.corners[j].y
		x, y := point.x, point.y

		intersects := ((yi > y) != (yj > y)) && (x < (xj-xi)*(y-yi)/(yj-yi)+xi)
		if intersects {
			inside = !inside
		}
	}
	return inside
}

func (bb *BoundingBox) AllCornersInside(corners []Position) bool {
	for _, corner := range corners {
		if !bb.IsPointInside(corner) {
			return false
		}
	}
	return true
}

func (bb *BoundingBox) EdgeIntersects(rect []Position) bool {
	for i, b1 := range bb.corners {
		j := (i + 1) % len(bb.corners)
		b2 := Position{x: bb.corners[j].x, y: bb.corners[j].y}

		for _, a1 := range rect {
			jj := (i + 1) % len(rect)
			a2 := rect[jj]

			if isEdgeIntersection(a1, a2, b1, b2) {
				return true
			}
		}
	}
	return false
}

func isEdgeIntersection(a1, a2, b1, b2 Position) bool {
	d1 := crossProduct(a2.Minus(&a1), b1.Minus(&a1))
	d2 := crossProduct(a2.Minus(&a1), b2.Minus(&a1))
	d3 := crossProduct(b2.Minus(&b1), a1.Minus(&b1))
	d4 := crossProduct(b2.Minus(&b1), a2.Minus(&b1))

	if (d1 > 0 && d2 < 0 || d1 < 0 && d2 > 0) &&
		(d3 > 0 && d4 < 0 || d3 < 0 && d4 > 0) {
		return true
	}

	return false
}

func crossProduct(p1, p2 Position) float64 {
	return p1.x*p2.y - p1.y*p2.x
}
