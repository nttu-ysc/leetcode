package problem1502

import "testing"

func TestCanMakeArithmeticProgression(t *testing.T) {
	tests := []struct {
		arr  []int
		want bool
	}{
		{
			arr:  []int{3, 5, 1},
			want: true,
		},
		{
			arr:  []int{1, 2, 4},
			want: false,
		},
	}

	for _, test := range tests {
		if got := canMakeArithmeticProgression(test.arr); got != test.want {
			t.Errorf("Test failed. The input arr=%v is expected to %t, but return %t", test.arr, test.want, got)
		}
	}
}
