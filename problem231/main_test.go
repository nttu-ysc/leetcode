package problem231

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{
			n:    1,
			want: true,
		},
		{
			n:    16,
			want: true,
		},
		{
			n:    3,
			want: false,
		},
		{
			n:    0,
			want: false,
		},
	}

	for _, test := range tests {
		if got := isPowerOfTwo(test.n); got != test.want {
			t.Errorf("Test failed. The input n = %d is expected to %t but return %t", test.n, test.want, got)
		}
	}
}
