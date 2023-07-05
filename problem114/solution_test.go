package problem114

import "testing"

func constructTreeNode(root []interface{}) *TreeNode {
	if len(root) == 0 {
		return nil
	}
	treeNode := &TreeNode{}

	q := []*TreeNode{}
	treeNode.Val = root[0].(int)
	root = root[1:]
	q = append(q, treeNode)

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		if len(root) > 0 {
			left := root[0]
			root = root[1:]
			if v, ok := left.(int); ok {
				cur.Left = &TreeNode{Val: v}
				q = append(q, cur.Left)
			}
		}
		if len(root) > 0 {
			right := root[0]
			root = root[1:]
			if v, ok := right.(int); ok {
				cur.Right = &TreeNode{Val: v}
				q = append(q, cur.Right)
			}
		}
	}

	return treeNode
}

func treeNode2Slice(root *TreeNode) []interface{} {
	if root == nil {
		return []interface{}{}
	}
	var ans []interface{}
	var q = []*TreeNode{root}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur != nil {
			ans = append(ans, cur.Val)
			q = append(q, cur.Left)
			q = append(q, cur.Right)
		} else {
			ans = append(ans, nil)
		}
	}
	return ans
}

func TestFlatten(t *testing.T) {
	testCases := []struct {
		root []interface{}
		want []interface{}
	}{
		{
			root: []interface{}{1, 2, 5, 3, 4, nil, 6},
			want: []interface{}{1, nil, 2, nil, 3, nil, 4, nil, 5, nil, 6},
		},
		{
			root: []interface{}{},
			want: []interface{}{},
		},
		{
			root: []interface{}{0},
			want: []interface{}{0},
		},
	}
	for _, tC := range testCases {
		root := constructTreeNode(tC.root)
		flatten(root)
		arr1 := treeNode2Slice(root)
		want := constructTreeNode(tC.want)
		arr2 := treeNode2Slice(want)

		if len(arr1) != len(arr2) {
			t.Errorf("Test failed.")
			continue
		}

		for i := 0; i < len(arr1); i++ {
			if arr1[i] != arr2[i] {
				t.Errorf("Test failed. Except to %v but return %v", arr1, arr2)
				continue
			}
		}
	}
}
