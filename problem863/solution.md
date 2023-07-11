# 863. All Nodes Distance K in Binary Tree

Given the root of a binary tree, the value of a target node target, and an integer k, return an array of the values of all nodes that have a distance k from the target node.

You can return the answer in any order.

## Explain

一般的 `TreeNode` 如果要找到距離 `root` 的點 `k` 非常容易只要使用 `DFS` 便能快速找到, 因此本題的關鍵點在於如何找到母節點

首先我們先對 `root` 遞迴便建立 `node` 與子節點, 使用 `haspMap` 來記錄

```go
func findParent(node *TreeNode, parent map[*TreeNode]*TreeNode) {
	if node == nil {
		return
	}
	if node.Left != nil {
		parent[node.Left] = node
	}
	findParent(node.Left, parent)
	if node.Right != nil {
		parent[node.Right] = node
	}
	findParent(node.Right, parent)
}
```

然後我們在 `DFS` 的時候再加入查找 `parent`, 並用 `visited` 來紀錄走訪過的點, 遍歷完所有的點即為所求
```go
func dfs(node *TreeNode, k int, parent map[*TreeNode]*TreeNode, visited map[*TreeNode]bool, ans *[]int) {
	if visited[node] {
		return
	}
	visited[node] = true
	if k == 0 {
		*ans = append(*ans, node.Val)
		return
	}
	if node.Left != nil {
		dfs(node.Left, k-1, parent, visited, ans)
	}
	if node.Right != nil {
		dfs(node.Right, k-1, parent, visited, ans)
	}
	if parent[node] != nil {
		dfs(parent[node], k-1, parent, visited, ans)
	}
}
```


## Code

[Code](./solution.go)