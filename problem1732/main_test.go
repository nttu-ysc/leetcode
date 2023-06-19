package problem1732

import "testing"

func TestLargestAltitude(t *testing.T) {
	tests := []struct {
		gain []int
		want int
	}{
		{
			gain: []int{-5, 1, 5, 0, -7},
			want: 1,
		},
		{
			gain: []int{-4, -3, -2, -1, 4, 3, 2},
			want: 0,
		},
	}

	for _, test := range tests {
		if got := largestAltitude(test.gain); got != test.want {
			t.Errorf("Test failed. The input gain = %v is excepted to %d but return %d", test.gain, test.want, got)
		}
	}
}
