package problem802

func eventualSafeNodes(graph [][]int) []int {
	ans := make([]int, 0)
	dp := make(map[int]bool)
	for i := 0; i < len(graph); i++ {
		if v, ok := (dp)[i]; ok {
			if v {
				ans = append(ans, i)
			}
			continue
		}
		for j := 0; j < len(graph[i]); j++ {
			visited := make([]int, len(graph))
			visited[i] = 1
			if !dfs(graph, graph[i][j], &visited, &dp) {
				goto b
			}
		}
		dp[i] = true
		ans = append(ans, i)
	b:
	}

	return ans
}

func dfs(graph [][]int, node int, visited *[]int, dp *map[int]bool) bool {
	if v, ok := (*dp)[node]; ok {
		return v
	}

	if (*visited)[node] == 1 {
		(*dp)[node] = false
		return false
	}
	(*visited)[node] = 1
	for i := 0; i < len(graph[node]); i++ {
		tmp := append([]int{}, *visited...)
		if !dfs(graph, graph[node][i], &tmp, dp) {
			(*dp)[graph[node][i]] = false
			return false
		} else {
			(*dp)[graph[node][i]] = true
		}
	}

	(*dp)[node] = true
	return true
}

func eventualSafeNodes2(graph [][]int) []int {
	visited, dfs_visited, cyclic := map[int]int{}, map[int]int{}, map[int]int{}
	res := []int{}

	for node := 0; node < len(graph); node++ {
		if visited[node] == 0 {
			detectCycle(node, graph, visited, dfs_visited, cyclic)
		}
	}
	for node := 0; node < len(graph); node++ {
		if cyclic[node] == 0 {
			res = append(res, node)
		}
	}
	return res
}

func detectCycle(node int, graph [][]int, visited map[int]int, dfs_visited map[int]int, cyclic map[int]int) bool {
	visited[node] = 1
	dfs_visited[node] = 1

	for _, adj := range graph[node] {
		if visited[adj] == 0 {
			if detectCycle(adj, graph, visited, dfs_visited, cyclic) {
				cyclic[node] = 1
				return true
			}
		}
		if dfs_visited[adj] == 1 {
			cyclic[node] = 1
			return true
		}
	}
	dfs_visited[node] = 0
	return false
}
