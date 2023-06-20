package problem242

import "testing"

func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		s, t string
		want bool
	}{
		{
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},

		{
			s:    "rat",
			t:    "car",
			want: false,
		},
		{
			s:    "aa",
			t:    "bb",
			want: false,
		},
	}

	for _, tC := range testCases {
		if got := isAnagram(tC.s, tC.t); got != tC.want {
			t.Errorf("Test failed. The input s=\"%s\", t = \"%s\" is expected to %t but return %t", tC.s, tC.t, tC.want, got)
		}
	}

}
