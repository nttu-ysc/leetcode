package problem852

import "testing"

func TestPeakIndexInMountainArray(t *testing.T) {
	testCases := []struct {
		arr  []int
		want int
	}{
		{
			arr:  []int{0, 1, 0},
			want: 1,
		},
		{
			arr:  []int{0, 2, 1, 0},
			want: 1,
		},
		{
			arr:  []int{0, 10, 5, 2},
			want: 1,
		},
		{
			arr:  []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19},
			want: 2,
		},
	}

	for _, tC := range testCases {
		if got := peakIndexInMountainArray(tC.arr); got != tC.want {
			t.Errorf("Test failed. The input arr = %v is expected to %d but return %d", tC.arr, tC.want, got)
		}
	}
}
