package problem200

var directions = [4][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func numIslands(grid [][]byte) int {
	var cnt int
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if dfs(grid, [2]int{row, col}) {
				cnt++
			}
		}
	}
	return cnt
}

func dfs(grid [][]byte, node [2]int) bool {
	row, col := node[0], node[1]
	if grid[row][col] == '0' {
		return false
	}
	grid[row][col] = '0'
	for _, dir := range directions {
		r, c := row+dir[0], col+dir[1]
		if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) {
			continue
		}

		if grid[r][c] == '1' {
			dfs(grid, [2]int{r, c})
		}
	}
	return true
}
