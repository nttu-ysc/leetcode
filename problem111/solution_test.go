package problem111

import "testing"

func constructTreeNode(arr []interface{}) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = arr[0].(int)
	q := make([]*TreeNode, 0)
	q = append(q, root)
	arr = arr[1:]
	for len(q) > 0 && len(arr) > 0 {
		cur := q[0]
		q = q[1:]
		left := arr[0]
		arr = arr[1:]
		if left != nil {
			cur.Left = &TreeNode{Val: left.(int)}
			q = append(q, cur.Left)
		}
		right := arr[0]
		arr = arr[1:]
		if right != nil {
			cur.Right = &TreeNode{Val: right.(int)}
			q = append(q, cur.Right)
		}
	}
	return root
}

func TestMinDepth(t *testing.T) {
	testCases := []struct {
		root []interface{}
		want int
	}{
		{
			root: []interface{}{3, 9, 20, nil, nil, 15, 7},
			want: 2,
		},
		{
			root: []interface{}{2, nil, 3, nil, 4, nil, 5, nil, 6},
			want: 5,
		},
		{
			root: []interface{}{},
			want: 0,
		},
	}

	for _, tC := range testCases {
		root := constructTreeNode(tC.root)
		if got := minDepth(root); got != tC.want {
			t.Errorf("Test failed. The input root %v is expected to %d but return %d", tC.root, tC.want, got)
		}
	}
}
