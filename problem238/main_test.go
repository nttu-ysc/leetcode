package problem238

import "testing"

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			nums: []int{-1, 1, 0, -3, 3},
			want: []int{0, 0, 9, 0, 0},
		},
	}

	for _, test := range tests {
		got := productExceptSelf(test.nums)

		if len(got) != len(test.nums) {
			t.Errorf("Test failed.")
		}

		for i := 0; i < len(got); i++ {
			if got[i] != test.want[i] {
				t.Errorf("Test failed. The input nums is expected to %v but return %v", test.want, got)
				break
			}
		}
	}
}
