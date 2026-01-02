package day4


func SolveP1(matrix [][]string) int {
	count := 0

	for row := range len(matrix) {
		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col] == "@" {
				if checkNeighbours(matrix, row, col) {
					count++
				}
			}
		}
	}

	return count
}

func checkNeighbours(matrix [][]string, row, col int) bool {
	count := 0
	// top left
	if row-1 >= 0 && col-1 >= 0 && matrix[row-1][col-1] == "@" {
		count++
	}

	// top middle
	if row-1 >= 0 && matrix[row-1][col] == "@" {
		count++
	}

	// top right
	if row-1 >= 0 && col+1 < len(matrix[row]) && matrix[row-1][col+1] == "@" {
		count++
	}

	// middle left
	if col-1 >= 0 && matrix[row][col-1] == "@" {
		count++
	}

	// middle right
	if col+1 < len(matrix[row]) && matrix[row][col+1] == "@" {
		count++
	}

	// bottom left
	if row+1 < len(matrix) && col-1 >= 0 && matrix[row+1][col-1] == "@" {
		count++
	}

	// bottom middle
	if row+1 < len(matrix) && matrix[row+1][col] == "@" {
		count++
	}

	// bottom right
	if row+1 < len(matrix) && col+1 < len(matrix[row]) && matrix[row+1][col+1] == "@" {
		count++
	}

	if count >= 4 {
		return false
	}

	return true
}

