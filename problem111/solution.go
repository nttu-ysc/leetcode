package problem111

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	var depth int
	for len(q) > 0 {
		n := len(q)
		depth++
		for i := 0; i < n; i++ {
			cur := q[0]
			q = q[1:]
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
		}
	}
	return depth
}
