package problem1970

var directions = [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func latestDayToCross(row int, col int, cells [][]int) int {
	g := make([][]int, row)
	for i := 0; i < row; i++ {
		g[i] = make([]int, col)
	}
	for i := 0; i < (row*col-1)/2; i++ {
		g[cells[i][0]-1][cells[i][1]-1] = 1
	}

	// binary search
	var l, r = 0, row*col - 1
	for l <= r {
		mid := l + (r-l)/2
		visited := make(map[[2]int]bool)
		for c := 0; c < col; c++ {
			if dfs(g, [2]int{0, c}, visited) {
				for i := mid; i < mid+1+(r-(mid+1))/2; i++ {
					g[cells[i][0]-1][cells[i][1]-1] = 1
				}
				l = mid + 1
				goto b
			}
		}
		for i := mid; i >= l+(mid-1-l)/2; i-- {
			g[cells[i][0]-1][cells[i][1]-1] = 0
		}
		r = mid - 1
	b:
	}
	return l - 1
}

func dfs(grid [][]int, node [2]int, visited map[[2]int]bool) bool {
	if grid[node[0]][node[1]] == 1 || visited[node] {
		return false
	}
	visited[node] = true
	if node[0] == len(grid)-1 {
		return true
	}
	for _, dir := range directions {
		x, y := node[0]+dir[0], node[1]+dir[1]
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
			continue
		}
		if dfs(grid, [2]int{x, y}, visited) {
			return true
		}
	}
	return false
}
