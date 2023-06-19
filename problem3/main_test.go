package problem3

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			s:    "abcabcbb",
			want: 3,
		},
		{
			s:    "bbbbb",
			want: 1,
		},
		{
			s:    "pwwkew",
			want: 3,
		},
		{
			s:    " ",
			want: 1,
		},
	}

	for _, test := range tests {
		if got := lengthOfLongestSubstring(test.s); got != test.want {
			t.Errorf("Test failed. The input s=\"%s\" is expected to %d but return %d", test.s, test.want, got)
		}
	}

}
