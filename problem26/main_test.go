package problem26

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 1, 2},
			want: 2,
		},
		{
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: 5,
		},
	}

	for _, tC := range testCases {
		if got := removeDuplicates2(tC.nums); got != tC.want {
			t.Errorf("Test failed. The nums = %v is expected to %d but return %d", tC.nums, tC.want, got)
		}
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	}
}

func BenchmarkRemoveDuplicates2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeDuplicates2([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	}
}
