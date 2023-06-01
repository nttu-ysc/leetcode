package problem226

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	dfs(root)

	return root
}

func dfs(root *TreeNode) {
	root.Left, root.Right = root.Right, root.Left

	if root.Left != nil {
		dfs(root.Left)
	}

	if root.Right != nil {
		dfs(root.Right)
	}
}
