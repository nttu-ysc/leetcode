package problem1091

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 || grid[len(grid)-1][len(grid[0])-1] == 1 {
		return -1
	}

	var dep = 0
	if bfs(grid, &dep, [][]int{{0, 0}}) {
		return dep
	}

	return -1
}

func bfs(grid [][]int, depth *int, current [][]int) bool {
	*depth++
	var next = [][]int{}
	for _, c := range current {
		if c[0] == len(grid)-1 && c[1] == len(grid[0])-1 {
			return true
		}
		process(grid, c[0], c[1]+1, &next)
		process(grid, c[0], c[1]-1, &next)
		process(grid, c[0]+1, c[1]+1, &next)
		process(grid, c[0]+1, c[1]-1, &next)
		process(grid, c[0]+1, c[1], &next)
		process(grid, c[0]-1, c[1]+1, &next)
		process(grid, c[0]-1, c[1]-1, &next)
		process(grid, c[0]-1, c[1], &next)
	}
	if len(next) == 0 {
		return false
	}

	return bfs(grid, depth, next)
}

func process(grid [][]int, i, j int, next *[][]int) {
	if i < 0 || j < 0 || i > len(grid)-1 || j > len(grid[0])-1 {
		return
	}

	if grid[i][j] == 0 {
		*next = append(*next, []int{i, j})
		grid[i][j] = 1
	}
}
