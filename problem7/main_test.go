package problem7

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{
			x:    123,
			want: 321,
		},
		{
			x:    -123,
			want: -321,
		},
		{
			x:    120,
			want: 21,
		},
		{
			x:    1534236469,
			want: 0,
		},
		{
			x:    1563847412,
			want: 0,
		},
	}

	for _, test := range tests {
		if got := reverse(test.x); got != test.want {
			t.Errorf("Test failed. The input x = %d is expected to %d but return %d", test.x, test.want, got)
		}
	}
}
