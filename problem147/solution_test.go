package problem147

import "testing"

func TestInsertionSortList(t *testing.T) {
	testCases := []struct {
		head []int
		want []int
	}{
		{
			head: []int{4, 2, 1, 3},
			want: []int{1, 2, 3, 4},
		},
		{
			head: []int{-1, 5, 3, 4, 0},
			want: []int{-1, 0, 3, 4, 5},
		},
	}

	for _, tC := range testCases {
		head := constructListNode(tC.head)
		want := constructListNode(tC.want)
		got := insertionSortList2(head)

		if want.Val != got.Val {
			t.Errorf("Test failed.")
			break
		}

		for want.Next != nil && got.Next != nil {
			want = want.Next
			got = got.Next
			if want.Val != got.Val {
				t.Errorf("Test failed.")
				break
			}
		}
	}
}

var head = constructListNode([]int{4, 2, 1, 3})

func BenchmarkInsertionSortList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertionSortList(head)
	}
}

func BenchmarkInsertionSortList2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertionSortList2(head)
	}
}
