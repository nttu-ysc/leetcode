package problem1161

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	levelSum := make(map[int]int, 0)
	dfs(root, &levelSum, 1)

	var level = 1

	for k := range levelSum {
		if levelSum[level] < levelSum[k] {
			level = k
		}
	}

	return level
}

func dfs(node *TreeNode, levelSum *map[int]int, level int) {
	if node == nil {
		return
	}

	(*levelSum)[level] += node.Val

	dfs(node.Left, levelSum, level+1)
	dfs(node.Right, levelSum, level+1)
}
