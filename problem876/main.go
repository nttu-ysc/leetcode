package problem876

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	var l []*ListNode
	l = append(l, head)
	for head.Next != nil {
		l = append(l, head.Next)
		head = head.Next
	}

	return l[len(l)/2]
}
