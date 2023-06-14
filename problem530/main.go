package problem530

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	return dfs(root, -2<<32, 2<<32-1)
}

func dfs(root *TreeNode, low, high int) int {
	if root == nil {
		return high - low
	}
	left := dfs(root.Left, low, root.Val)
	right := dfs(root.Right, root.Val, high)
	return min(left, right)
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
