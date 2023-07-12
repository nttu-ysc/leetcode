# 802. Find Eventual Safe States

There is a directed graph of n nodes with each node labeled from `0` to `n - 1`. The graph is represented by a 0-indexed 2D integer array `graph` where `graph[i]` is an integer array of nodes adjacent to node `i`, meaning there is an edge from node `i` to each node in `graph[i]`.

A node is a **terminal node** if there are no outgoing edges. A node is a **safe node** if every possible path starting from that node leads to a **terminal node** (or another safe node).

Return an array containing all the **safe nodes** of the graph. The answer should be sorted in **ascending** order.

## Explain

該題目就是想要找到一個不會形成迴圈的 `node`, 因此我們可以寫一個 `func` 來確認該 `node` 是否有迴圈

這邊會用兩個 `map` 來紀錄

- `visited` 為 `1` 表示已經走訪過此 `node`
- `dfs_visited` 為 `1` 表示此節點有迴圈

```go
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
```


## Code

[Solution](./method2.go)