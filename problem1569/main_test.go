package problem1569

import "testing"

func TestNumOfWays(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{2, 1, 3},
			want: 1,
		},
		{
			nums: []int{3, 4, 5, 1, 2},
			want: 5,
		},
		{
			nums: []int{1, 2, 3},
			want: 0,
		},
		{
			nums: []int{31, 23, 14, 24, 15, 12, 25, 28, 5, 35, 17, 6, 9, 11, 1, 27, 18, 20, 2, 3, 33, 10, 13, 4, 7, 36, 32, 29, 8, 30, 26, 19, 34, 22, 21, 16},
			want: 936157466,
		},
		{
			nums: []int{10, 23, 12, 18, 4, 29, 2, 8, 41, 31, 25, 21, 14, 35, 26, 5, 19, 43, 22, 37, 9, 20, 44, 28, 1, 39, 30, 38, 36, 6, 13, 16, 27, 17, 34, 7, 15, 3, 11, 24, 42, 33, 40, 32},
			want: 182440977,
		},
		{
			nums: []int{30, 11, 46, 3, 29, 22, 37, 32, 13, 49, 48, 16, 40, 8, 24, 44, 9, 39, 25, 21, 41, 26, 43, 19, 47, 7, 10, 31, 45, 4, 35, 14, 20, 23, 15, 17, 38, 2, 6, 18, 5, 33, 27, 36, 42, 28, 12, 34, 1},
			want: 137401437,
		},
		{
			nums: []int{74, 24, 70, 11, 6, 4, 59, 9, 36, 82, 80, 30, 46, 31, 22, 34, 8, 69, 32, 57, 18, 21, 37, 83, 55, 38, 41, 72, 48, 65, 27, 60, 73, 58, 68, 50, 16, 77, 75, 20, 81, 3, 61, 13, 10, 29, 62, 49, 12, 66, 39, 45, 28, 40, 42, 52, 78, 56, 44, 17, 14, 67, 35, 26, 19, 5, 63, 51, 43, 23, 79, 2, 54, 47, 76, 53, 7, 25, 64, 33, 1, 15, 71},
			want: 901891111,
		},
	}

	for _, test := range tests {
		t.Logf("start...")
		if got := numOfWays(test.nums); got != test.want {
			t.Errorf("Test failed. The input nums=%v is expected to %d but return %d", test.nums, test.want, got)
		}
	}
}
