package problem2328

var directions = [4][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

const MOD = 1e9 + 7

func countPaths(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	dp := make([][]int, rows)
	for k := range dp {
		dp[k] = make([]int, cols)
	}
	var ans int

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			ans = (ans + dfs(grid, row, col, &dp)) % MOD
		}
	}

	return ans % MOD
}

func dfs(grid [][]int, row, col int, dp *[][]int) int {
	rows, cols := len(grid), len(grid[0])
	if (*dp)[row][col] > 0 {
		return (*dp)[row][col]
	}

	ans := 1
	for _, d := range directions {
		r, c := row+d[0], col+d[1]
		if r < 0 || c < 0 || r >= rows || c >= cols {
			continue
		}

		if grid[r][c] > grid[row][col] {
			ans = (ans + dfs(grid, r, c, dp)) % MOD
		}
	}

	(*dp)[row][col] = ans

	return (*dp)[row][col]
}
