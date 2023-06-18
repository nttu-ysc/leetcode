package problem329

var directions = [][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func longestIncreasingPath(matrix [][]int) int {
	rows, cols := len(matrix), len(matrix[0])
	dp := make([][]int, rows)
	for k := range dp {
		dp[k] = make([]int, cols)
	}
	var ans int

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			ans = max(dfs(matrix, row, col, &dp), ans)
		}
	}

	return ans
}

func dfs(matrix [][]int, row, col int, dp *[][]int) int {
	rows, cols := len(matrix), len(matrix[0])
	if (*dp)[row][col] > 0 {
		return (*dp)[row][col]
	}

	ans := 1

	for _, dir := range directions {
		r, c := row+dir[0], col+dir[1]
		if r < 0 || c < 0 || r >= rows || c >= cols {
			continue
		}

		if matrix[r][c] > matrix[row][col] {
			ans = max(dfs(matrix, r, c, dp)+1, ans)
		}
	}
	(*dp)[row][col] = ans
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
