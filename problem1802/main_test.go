package problem1802

import "testing"

func TestMaxValue(t *testing.T) {
	tests := []struct {
		n      int
		index  int
		maxSum int
		want   int
	}{
		{
			n:      4,
			index:  2,
			maxSum: 6,
			want:   2,
		},
		{
			n:      6,
			index:  1,
			maxSum: 10,
			want:   3,
		},
	}

	for _, test := range tests {
		if got := maxValue(test.n, test.index, test.maxSum); got != test.want {
			t.Errorf("Test failed. The n=%d index=%d maxSum=%d is excepted tot %d but return %d", test.n, test.index, test.maxSum, test.want, got)
		}
	}
}
