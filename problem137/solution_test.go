package problem137

import "testing"

func TestSingleNumber(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{2, 2, 3, 2},
			want: 3,
		},
		{
			nums: []int{0, 1, 0, 1, 0, 1, 99},
			want: 99,
		},
		{
			nums: []int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2},
			want: -4,
		},
	}

	for _, tC := range testCases {
		if got := singleNumber2(tC.nums); got != tC.want {
			t.Errorf("Test failed. The input nums = %v is expected to %d but return %d", tC.nums, tC.want, got)
		}
	}
}

var nums = []int{0, 1, 0, 1, 0, 1, 99}

func BenchmarkSingleNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		singleNumber(nums)
	}
}

func BenchmarkSingleNumber2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		singleNumber2(nums)
	}
}
