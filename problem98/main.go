package problem98

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	l := new([]int)
	return dfs(root, l)

}

func dfs(root *TreeNode, l *[]int) bool {
	if root.Left != nil {
		if !dfs(root.Left, l) {
			return false
		}
	}
	if len(*l) > 0 && root.Val <= (*l)[len(*l)-1] {
		return false
	}
	*l = append(*l, root.Val)
	if root.Right != nil {
		if !dfs(root.Right, l) {
			return false
		}
	}

	return true
}

func construct(arr []interface{}) *TreeNode {
	root := new(TreeNode)

	root.Val = arr[0].(int)

	var q []*TreeNode
	q = append(q, root)

	for i := 1; i < len(arr); i += 2 {
		tmp := q[0]
		q = q[1:]

		if cur, ok := arr[i].(int); !ok {
			tmp.Left = nil
		} else {
			tmp.Left = &TreeNode{Val: cur}
			q = append(q, tmp.Left)
		}

		if cur, ok := arr[i+1].(int); !ok {
			tmp.Right = nil
		} else {
			tmp.Right = &TreeNode{Val: cur}
			q = append(q, tmp.Right)
		}
	}

	return root
}
