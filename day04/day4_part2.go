package day4

func SolveP2(matrix [][]string) int {
	count := 0
	previous := -1
	modified := make([][]string, len(matrix))
	for i := range modified {
		modified[i] = make([]string, len(matrix[0]))
	}

	for count != previous {
		previous = count
		run(matrix, modified, &count)
		postProccess(modified)
		matrix = modified
	}

	return count
}

func run(matrix [][]string, modified [][]string, count *int) {
	for row := range len(matrix) {
		for col := 0; col < len(matrix[row]); col++ {
			modified[row][col] = matrix[row][col]
			if matrix[row][col] == "@" {
				if checkNeighbours(matrix, row, col) {
					modified[row][col] = "x"
					*count++
				}
			}
		}
	}
}

func postProccess(matrix [][]string) [][]string {
	for row := range len(matrix) {
		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col] == "x" {
				matrix[row][col] = "."
			}
		}
	}

	return matrix
}
