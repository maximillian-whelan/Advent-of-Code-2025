package day9


func LargestArea(fp string) int {
	positions := readFile(fp).corners
	largest := 0
	for i := range len(positions) {
		for j := range len(positions) {
			if i != j {
				rect := Rectangle{positions[i], positions[j]}
				area := rect.Area()

				if area > largest {
					largest = area
				}
			}
		}
	}

	return largest
}

