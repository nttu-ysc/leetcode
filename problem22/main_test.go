package problem22

import "testing"

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		n    int
		want []string
	}{
		{
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			n:    1,
			want: []string{"()"},
		},
	}

	for _, test := range tests {
		got := generateParenthesis(test.n)
		if len(got) != len(test.want) {
			t.Errorf("Test failed.")
		}

		for i := 0; i < len(got); i++ {
			if got[i] != test.want[i] {
				t.Errorf("Test failed. The input n = %d is expected to %v but return %v", test.n, test.want, got)
				break
			}
		}
	}
}
