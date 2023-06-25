package problem1544

import "testing"

func TestMakeGood(t *testing.T) {
	testCases := []struct {
		s    string
		want string
	}{
		{
			s:    "leEeetcode",
			want: "leetcode",
		},
		{
			s:    "abBAcC",
			want: "",
		},
	}

	for _, tC := range testCases {
		if got := makeGood(tC.s); got != tC.want {
			t.Errorf("Test failed. The s = \"%s\" is expected to \"%s\" but return \"%s\"", tC.s, tC.want, got)
		}
	}
}
