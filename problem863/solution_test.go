package problem863

import "testing"

func constructTreeNode(arr []interface{}, target int) (*TreeNode, *TreeNode) {
	root := new(TreeNode)
	root.Val = arr[0].(int)
	arr = arr[1:]
	var q []*TreeNode
	q = append(q, root)
	var t *TreeNode

	for len(q) > 0 && len(arr) > 0 {
		cur := q[0]
		q = q[1:]

		left := arr[0]
		arr = arr[1:]
		if left != nil {
			cur.Left = &TreeNode{Val: left.(int)}
			q = append(q, cur.Left)
			if cur.Left.Val == target {
				t = cur.Left
			}
		}

		right := arr[0]
		arr = arr[1:]
		if right != nil {
			cur.Right = &TreeNode{Val: right.(int)}
			q = append(q, cur.Right)
			if cur.Right.Val == target {
				t = cur.Right
			}
		}
	}
	if t == nil {
		t = &TreeNode{Val: target}
	}
	return root, t
}

func TestDistanceK(t *testing.T) {
	testCase := []struct {
		root   []interface{}
		target int
		k      int
		want   []int
	}{
		{
			root:   []interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4},
			target: 5,
			k:      2,
			want:   []int{7, 4, 1},
		},
		{
			root:   []interface{}{1},
			target: 1,
			k:      3,
			want:   []int{},
		},
	}

	for _, tC := range testCase {
		root, target := constructTreeNode(tC.root, tC.target)
		got := distanceK(root, target, tC.k)

		if len(got) != len(tC.want) {
			t.Errorf("Test failed.")
		}

		for i := 0; i < len(got); i++ {
			if got[i] != tC.want[i] {
				t.Errorf("Test failed. The input is expected to %v but return %v", tC.want, got)
				break
			}
		}
	}
}
