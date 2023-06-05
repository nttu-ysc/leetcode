package problem785

func isBipartite(graph [][]int) bool {
	color := make([]int, len(graph))
	for k := range color {
		color[k] = -1
	}

	for i := 0; i < len(color); i++ {
		if color[i] == -1 {
			if !dfs(graph, i, color, 0) {
				return false
			}
		}
	}

	return true
}

func dfs(graph [][]int, node int, color []int, criteria int) bool {
	if color[node] != -1 {
		return color[node] == criteria
	}

	color[node] = criteria
	next := 0
	if criteria == next {
		next = 1
	}

	for i := 0; i < len(graph[node]); i++ {
		if !dfs(graph, graph[node][i], color, next) {
			return false
		}
	}

	return true
}
