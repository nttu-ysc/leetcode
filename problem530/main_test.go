package problem530

import "testing"

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

func TestMinimumDifference(t *testing.T) {
	tests := []struct {
		root []interface{}
		want int
	}{
		{
			root: []interface{}{4, 2, 6, 1, 3},
			want: 1,
		},
		{
			root: []interface{}{1, 0, 48, nil, nil, 12, 49},
			want: 1,
		},
	}

	for _, test := range tests {
		root := construct(test.root)
		if got := getMinimumDifference(root); got != test.want {
			t.Errorf("Test failed. The input root=%v is excepted to %d but return %d", test.root, test.want, got)
		}
	}
}
