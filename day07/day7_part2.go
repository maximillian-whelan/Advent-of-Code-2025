package day7

type StartPos struct {
	row int
	col int
}

func NumberOfTimelines(fp string) int {

	manifold := ReadManifold(fp)
	sp := &StartPos{0, 0}

	for row := range len(manifold) {
		for col := range len(manifold[row]) {
			if manifold[row][col] == 'S' {
				sp = &StartPos{row, col}
			}
		}
	}

	memo := make([][]int, len(manifold))
	for i := range memo {
		memo[i] = make([]int, len(manifold[0]))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return dfs(manifold, memo, sp.row, sp.col)
}

func dfs(manifold [][]byte, memo [][]int, row, col int) int {

	// if we go out of bounds early return with none found or 0
	if row < 0 || col < 0 || row >= len(manifold) || col >= len(manifold[col]) {
		return 0
	}

	// base case
	// we are at the bottom so we have completed a route
	if row == len(manifold)-1 {
		return 1
	}

	count := 0
	// dfs was too slow so had to introduce a memo to keep track of pre-solved paths
	// would be more difficult if we can go up
	if memo[row][col] != -1 {
		return memo[row][col]
	}

	if manifold[row][col] == '^' {
		count += dfs(manifold, memo, row+1, col-1) + dfs(manifold, memo, row+1, col+1)
	} else {
		count += dfs(manifold, memo, row+1, col)
	}

	memo[row][col] = count
	return count
}
