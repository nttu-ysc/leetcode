package problem14

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs []string
		want string
	}{
		{
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			strs: []string{"ab", "a"},
			want: "a",
		},
		{
			strs: []string{"aaa", "aa", "aaa"},
			want: "aa",
		},
	}

	for _, test := range tests {
		if got := longestCommonPrefix(test.strs); got != test.want {
			t.Errorf("Test failed. the input strs=%v is excepted to \"%s\" but return \"%s\"", test.strs, test.want, got)
		}
	}

}
