package problem104

import "testing"

func constructTreeNode(arr []interface{}) *TreeNode {
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

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		root []interface{}
		want int
	}{
		{
			root: []interface{}{3, 9, 20, nil, nil, 15, 7},
			want: 3,
		},
		{
			root: []interface{}{1, nil, 2},
			want: 2,
		},
	}

	for _, test := range tests {
		root := constructTreeNode(test.root)
		if got := maxDepth(root); got != test.want {
			t.Errorf("Test failed. The root = %v is expected to %d but return %d", test.root, test.want, got)
		}
	}
}
