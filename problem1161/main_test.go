package problem1161

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

func TestMaxLevelSum(t *testing.T) {
	tests := []struct {
		root []interface{}
		want int
	}{
		{
			root: []interface{}{1, 7, 0, 7, -8, nil, nil},
			want: 2,
		},
		{
			root: []interface{}{989, nil, 10250, 98693, -89388, nil, nil, nil, -32127},
			want: 2,
		},
		{
			root: []interface{}{-100, -200, -300, -20, -5, -10, nil},
			want: 3,
		},
	}

	for _, test := range tests {
		root := construct(test.root)
		if got := maxLevelSum(root); got != test.want {
			t.Errorf("Test failed. The input root=%v is excepted to %d but return %d", test.root, test.want, got)
		}
	}
}

func BenchmarkMaxLevelSumDfs(b *testing.B) {
	root := construct([]interface{}{989, nil, 10250, 98693, -89388, nil, nil, nil, -32127})
	for i := 0; i < b.N; i++ {
		maxLevelSum(root)
	}
}

func BenchmarkMaxLevelSumBfs(b *testing.B) {
	root := construct([]interface{}{989, nil, 10250, 98693, -89388, nil, nil, nil, -32127})
	for i := 0; i < b.N; i++ {
		maxLevelSum2(root)
	}
}
