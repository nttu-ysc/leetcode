package problem802

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
