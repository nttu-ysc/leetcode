package problem864

import "fmt"

func shortestPathAllKeys(grid []string) int {
	// ↓→↑←
	var directions = [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	m, n := len(grid), len(grid[0])
	var keyCount int
	var q [][3]int
	var visited = make(map[string]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] >= 'a' && grid[i][j] <= 'f' {
				keyCount++
			} else if grid[i][j] == '@' {
				q = append(q, [3]int{i, j, 0})
				visited[fmt.Sprintf("%d_%d_%d", i, j, 0)]++
			}
		}
	}

	var step int
	for len(q) > 0 {
		for i := len(q); i > 0; i-- {
			curX, curY, curKeys := q[0][0], q[0][1], q[0][2]
			if curKeys == (1<<keyCount)-1 {
				return step
			}
			q = q[1:]
			for _, dir := range directions {
				x := curX + dir[0]
				y := curY + dir[1]
				keys := curKeys
				if x < 0 || x >= m || y < 0 || y >= n {
					continue
				}
				c := grid[x][y]
				// meet wall
				if c == '#' {
					continue
				}
				// no corresponding keys
				if c >= 'A' && c <= 'F' && ((keys >> (c - 'A') & 1) == 0) {
					continue
				}
				if c >= 'a' && c <= 'f' {
					keys |= 1 << (c - 'a')
				}
				hashKey := fmt.Sprintf("%d_%d_%d", x, y, keys)
				if visited[hashKey] <= 0 {
					visited[hashKey]++
					q = append(q, [3]int{x, y, keys})
				}
			}
		}
		step++
	}
	return -1
}
