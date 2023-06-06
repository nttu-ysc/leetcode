package problem876

import (
	"testing"
)

func TestMiddleNode(t *testing.T) {
	tests := []struct {
		head []int
		want []int
	}{
		{
			head: []int{1, 2, 3, 4, 5},
			want: []int{3, 4, 5},
		},
		{
			head: []int{1, 2, 3, 4, 5, 6},
			want: []int{4, 5, 6},
		},
	}

	for _, test := range tests {
		ln := constructor(test.head)
		mid := constructor(test.want)
		got := middleNode(ln)

		wantArr := []int{mid.Val}
		gotArr := []int{got.Val}

		for mid.Next != nil {
			wantArr = append(wantArr, mid.Next.Val)
			mid = mid.Next
		}

		for got.Next != nil {
			gotArr = append(gotArr, got.Next.Val)
			got = got.Next
		}

		if len(gotArr) != len(wantArr) {
			t.Errorf("Test failed.")
			break
		}

		for i := 0; i < len(gotArr); i++ {
			if gotArr[i] != wantArr[i] {
				t.Errorf("Test failed.")
				break
			}
		}
	}
}

func constructor(head []int) *ListNode {
	if len(head) == 0 {
		return nil
	}

	l := new(ListNode)
	l.Val = head[0]
	l.Next = constructor(head[1:])
	return l
}
