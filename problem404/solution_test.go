package problem404

import (
	"testing"
)

func constructTreeNode(root []interface{}) *TreeNode {
	treeNode := &TreeNode{Val: root[0].(int)}
	q := make([]*TreeNode, 0)
	q = append(q, treeNode)
	root = root[1:]
	for len(q) > 0 && len(root) > 1 {
		cur := q[0]
		q = q[1:]
		if root[0] != nil {
			cur.Left = &TreeNode{Val: root[0].(int)}
			q = append(q, cur.Left)
		}
		root = root[1:]
		if root[0] != nil {
			cur.Right = &TreeNode{Val: root[0].(int)}
			q = append(q, cur.Right)
		}
		root = root[1:]
	}
	return treeNode
}

func TestSumOfLeftLeaves(t *testing.T) {
	testCases := []struct {
		root *TreeNode
		want int
	}{
		{
			root: constructTreeNode([]interface{}{3, 9, 20, nil, nil, 15, 7}),
			want: 24,
		},
		{
			root: constructTreeNode([]interface{}{1}),
			want: 0,
		},
	}

	for _, tC := range testCases {
		if got := sumOfLeftLeaves(tC.root); got != tC.want {
			t.Errorf("Test failed. The input root = %v is expected to %d but return %d", tC.root, tC.want, got)
		}
	}
}
