package problem147

func insertionSortList2(head *ListNode) *ListNode {
	dummy := new(ListNode)
	for head != nil {
		cur := dummy
		for ; cur.Next != nil && cur.Next.Val < head.Val; cur = cur.Next {
		}
		cur.Next, head.Next, head = head, cur.Next, head.Next
	}
	return dummy.Next
}
