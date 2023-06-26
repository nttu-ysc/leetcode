package problem147

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func constructListNode(arr []int) *ListNode {
	if len(arr) <= 0 {
		return nil
	}
	head := new(ListNode)
	head.Val = arr[0]
	head.Next = constructListNode(arr[1:])
	return head
}

func insertionSortList(head *ListNode) *ListNode {
	arr := make([]int, 0)
	for i := head; i != nil; i = i.Next {
		arr = append(arr, i.Val)
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return constructListNode(arr)
}
