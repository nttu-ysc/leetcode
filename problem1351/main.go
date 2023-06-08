package problem1351

func countNegatives(grid [][]int) int {
	board := len(grid[0])
	cnt := len(grid) * board
	for i := 0; i < len(grid); i++ {
		for j := 0; j < board; j++ {
			if grid[i][j] >= 0 {
				cnt--
			} else {
				board = j
			}
		}
	}

	return cnt
}
