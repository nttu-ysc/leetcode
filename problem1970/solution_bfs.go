package problem1970

func latestDayToCross2(row, col int, cells [][]int) int {
	var l, r = col - 1, row * col
	for l < r {
		mid := (l + r) >> 1
		if bfs(mid+1, row, col, cells) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func bfs(day, row, col int, cells [][]int) bool {
	grid := make([][]int, row)
	for i := 0; i < row; i++ {
		grid[i] = make([]int, col)
	}
	for i := 0; i < day; i++ {
		grid[cells[i][0]-1][cells[i][1]-1] = 1
	}
	q := [][2]int{}
	for c := 0; c < col; c++ {
		if grid[0][c] == 0 {
			grid[0][c] = 1
			q = append(q, [2]int{0, c})
		}
	}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur[0] == row-1 {
			return true
		}
		for _, dir := range directions {
			x, y := cur[0]+dir[0], cur[1]+dir[1]
			if x < 0 || x >= row || y < 0 || y >= col || grid[x][y] == 1 {
				continue
			}
			grid[x][y] = 1
			q = append(q, [2]int{x, y})
		}
	}
	return false
}
