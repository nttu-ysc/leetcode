package problem21

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	tmp := new(ListNode)
	ans := tmp

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tmp.Next = &ListNode{Val: list1.Val}
			tmp = tmp.Next
			list1 = list1.Next
		} else {
			tmp.Next = &ListNode{Val: list2.Val}
			tmp = tmp.Next
			list2 = list2.Next
		}
	}
	for list1 != nil {
		tmp.Next = &ListNode{Val: list1.Val}
		tmp = tmp.Next
		list1 = list1.Next
	}
	for list2 != nil {
		tmp.Next = &ListNode{Val: list2.Val}
		tmp = tmp.Next
		list2 = list2.Next
	}

	return ans.Next
}
