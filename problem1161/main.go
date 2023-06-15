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

func maxLevelSum2(root *TreeNode) int {
	queue := []*TreeNode{root}
	level := 0
	index := -1
	diff := -1 << 32

	for len(queue) > 0 {
		level++
		var sum int
		l := len(queue)
		for i := 0; i < l; i++ {
			temp := queue[0]
			queue = queue[1:]

			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}

			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
			sum += temp.Val
		}

		if sum > diff {
			diff = sum
			index = level
		}
	}

	return index
}
