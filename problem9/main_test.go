package problem9

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		x    int
		want bool
	}{
		{
			x:    121,
			want: true,
		},
		{
			x:    -121,
			want: false,
		},
		{
			x:    10,
			want: false,
		},
		{
			x:    1001,
			want: true,
		},
	}

	for _, test := range tests {
		if got := isPalindrome(test.x); got != test.want {
			t.Errorf("Test failed. the input x=%d is excepted to %t but return %t", test.x, test.want, got)
		}
	}
}
