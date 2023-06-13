package problem2352

func equalPairs(grid [][]int) int {
	n := len(grid)
	var count int
	for i := 0; i < n; i++ {
		row := grid[i]

		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if grid[k][j] != row[k] {
					goto b
				}
			}
			count++
		b:
		}
	}

	return count
}
