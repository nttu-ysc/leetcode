package problem10

import "testing"

func TestIsMatch(t *testing.T) {
	testCases := []struct {
		s    string
		p    string
		want bool
	}{
		{
			s:    "aa",
			p:    "a",
			want: false,
		},
		{
			s:    "aa",
			p:    "a*",
			want: true,
		},
		{
			s:    "ab",
			p:    ".*",
			want: true,
		},
	}

	for _, tC := range testCases {
		if got := isMatch(tC.s, tC.p); got != tC.want {
			t.Errorf("Test failed. The input s = \"%s\", p = \"%s\" is expected to %t but return %t", tC.s, tC.p, tC.want, got)
		}
	}
}
