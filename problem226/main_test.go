package problem226

import "testing"

func TestInvertTree(t *testing.T) {
	tests := []struct {
		root []int
		want []int
	}{
		{
			root: []int{4, 2, 7, 1, 3, 6, 9},
			want: []int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			root: []int{2, 1, 3},
			want: []int{2, 3, 1},
		},
		{
			root: []int{},
			want: []int{},
		},
	}

	for _, test := range tests {
		tree := construct(test.root)
		gotTree := invertTree(tree)
		got := []int{}
		bfs(gotTree, &got)
		for i := 0; i < min(len(got), len(test.want)); i++ {
			if got[i] != test.want[i] {
				t.Errorf("Test failed. the input root=%v is excepted to %v but return %v", test.
					root, test.want, got)
				break
			}
		}
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func bfs(root *TreeNode, l *[]int) {
	q := []*TreeNode{}
	if root != nil {
		q = append(q, root)
	}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		*l = append(*l, cur.Val)

		if cur.Left != nil {
			q = append(q, cur.Left)
		}

		if cur.Right != nil {
			q = append(q, cur.Right)
		}
	}
}

func construct(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	root := new(TreeNode)

	root.Val = arr[0]

	var q []*TreeNode
	q = append(q, root)

	for i := 1; i < len(arr); i += 2 {
		tmp := q[0]
		q = q[1:]

		cur := arr[i]
		tmp.Left = &TreeNode{Val: cur}
		q = append(q, tmp.Left)

		cur = arr[i+1]
		tmp.Right = &TreeNode{Val: cur}
		q = append(q, tmp.Right)
	}

	return root
}
