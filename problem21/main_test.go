package problem21

import "testing"

func constructListNode(list []int) *ListNode {
	var tmp *ListNode
	for i := len(list) - 1; i >= 0; i-- {
		tmp = &ListNode{
			Val:  list[i],
			Next: tmp,
		}
	}

	return tmp
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		list1 ListNode
		list2 ListNode
		want  []int
	}{
		{
			list1: *constructListNode([]int{1, 2, 4}),
			list2: *constructListNode([]int{1, 3, 4}),
			want:  []int{1, 1, 2, 3, 4, 4},
		},
	}

	for _, test := range tests {
		gotList := mergeTwoLists(&test.list1, &test.list2)
		got := []int{gotList.Val}
		for gotList.Next != nil {
			got = append(got, gotList.Next.Val)
			gotList = gotList.Next
		}

		if len(test.want) != len(got) {
			t.Errorf("Test failed.")
		}

		for i := 0; i < len(got); i++ {
			if got[i] != test.want[i] {
				t.Errorf("Test failed. The input list1 = %v, list2 = %v is expected to %v but return %v", test.list1, test.list2, test.want, got)
			}
		}
	}
}
