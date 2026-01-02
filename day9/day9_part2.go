package day9

// breakdown of task:

// I need to check each corner agaisnt other corners
// the corners must be adjacent so no point checking corners on the same column or row (slight optimisation here)
// each corner must be within the bouding box

// I need to create a bounding box for the search area
// bounding box is defined by the # locations (corners) or the locations or my input

// if any of the points inside my rectangle are outside of the bounding box or search area
// this rectangle can be ignored
// not sure on how to check if rectangle is inside or outside of bounding box

func MakeLargestRectangle(fp string) int {
	bb := readFile(fp)
	largest := 0

	for i := range len(bb.corners) {
		for j := range len(bb.corners) {
			if i != j ||
				bb.corners[i].x != bb.corners[j].x ||
				bb.corners[i].y != bb.corners[j].y {

				rect := Rectangle{
					start: bb.corners[i],
					end:   bb.corners[j],
				}

				if bb.RectInsideBoundingBox(rect) {
					area := rect.Area()
					if area > largest {
						largest = area
					}
				}
			}
		}
	}

	return largest
}
