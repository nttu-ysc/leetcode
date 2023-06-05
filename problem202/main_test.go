package problem202

import "testing"

func TestIsHappy(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{
			n:    19,
			want: true,
		},
		{
			n:    2,
			want: false,
		},
	}

	for _, test := range tests {
		if got := isHappy(test.n); test.want != got {
			t.Errorf("Test failed. The input n=%d is excepted to %t but return %t", test.n, test.want, got)
		}
	}
}
