package problem543

import (
	"testing"
)

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

		if i+1 > len(arr)-1 {
			break
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

func TestDiameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		root []interface{}
		want int
	}{
		{
			root: []interface{}{1, 2, 3, 4, 5},
			want: 3,
		},
		{
			root: []interface{}{1, 2},
			want: 1,
		},
		{
			root: []interface{}{4, -7, -3, nil, nil, -9, -3, 9, -7, -4, nil, 6, nil, -6, -6, nil, nil, 0, 6, 5, nil, 9, nil, nil, -1, -4, nil, nil, nil, -2},
			want: 8,
		},
	}

	for _, test := range tests {
		root := construct(test.root)
		if got := diameterOfBinaryTree(root); got != test.want {
			t.Errorf("Test failed. The input root = %v is excepted to %d but return %d", test.root, test.want, got)
		}
	}
}
