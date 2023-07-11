package problem863

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	var ans []int
	if root == nil {
		return ans
	}
	parent := make(map[*TreeNode]*TreeNode)
	visited := make(map[*TreeNode]bool)
	findParent(root, parent)
	dfs(target, k, parent, visited, &ans)
	return ans
}

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
