package problem2101

import "math"

func maximumDetonation(bombs [][]int) int {
	direct := make([][]int, len(bombs))
	for i, v0 := range bombs {
		for j, v1 := range bombs {
			if i == j {
				continue
			}
			if distanceInner(v0[:2], v1[:2], v0[2]) {
				direct[i] = append(direct[i], j)
			}
		}
	}

	var ans int
	for i := range direct {
		var count int
		dfs(direct, i, make([]bool, len(bombs)), &count)
		if count > ans {
			ans = count
		}
	}

	return ans
}

func distanceInner(a, b []int, d int) bool {
	return math.Pow(float64(d), 2) >= math.Pow(float64(a[0]-b[0]), 2)+math.Pow(float64(a[1]-b[1]), 2)
}

func dfs(direct [][]int, index int, visited []bool, count *int) {
	if visited[index] {
		return
	}

	visited[index] = true
	(*count)++
	for _, bomb := range direct[index] {
		dfs(direct, bomb, visited, count)
	}
}
