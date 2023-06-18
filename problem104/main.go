package problem104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return dfs(root)
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := dfs(node.Left)
	right := dfs(node.Right)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
