package problem543

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var sum int
	dfs(root, &sum)
	return sum
}

func dfs(node *TreeNode, sum *int) int {
	if node == nil {
		return 0
	}
	var left, right int
	left = dfs(node.Left, sum)
	right = dfs(node.Right, sum)
	*sum = max(*sum, left+right)
	return max(left, right) + 1
}

func max(arr ...int) int {
	if len(arr) < 1 {
		return -1
	}
	ans := arr[0]
	for _, v := range arr {
		if v > ans {
			ans = v
		}
	}

	return ans
}
