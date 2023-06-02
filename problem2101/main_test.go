package problem2101

import "testing"

func TestMaximumDetonation(t *testing.T) {
	tests := []struct {
		bombs [][]int
		want  int
	}{
		{
			bombs: [][]int{
				{2, 1, 3},
				{6, 1, 4},
			},
			want: 2,
		},
		{
			bombs: [][]int{
				{1, 1, 5},
				{10, 10, 5},
			},
			want: 1,
		},
		{
			bombs: [][]int{
				{1, 2, 3},
				{2, 3, 1},
				{3, 4, 2},
				{4, 5, 3},
				{5, 6, 4},
			},
			want: 5,
		},
	}

	for _, test := range tests {
		if got := maximumDetonation(test.bombs); got != test.want {
			t.Errorf("Test failed. The intput bombs=%v is excepted to %d but return %d", test.bombs, test.want, got)
		}
	}
}
