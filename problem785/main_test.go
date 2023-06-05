package problem785

import "testing"

func TestIsBipartite(t *testing.T) {
	tests := []struct {
		graph [][]int
		want  bool
	}{
		{
			graph: [][]int{
				{1, 2, 3},
				{0, 2},
				{0, 1, 3},
				{0, 2},
			},
			want: false,
		},
		{
			graph: [][]int{
				{1, 3},
				{0, 2},
				{1, 3},
				{0, 2},
			},
			want: true,
		},
	}

	for _, test := range tests {
		if got := isBipartite(test.graph); got != test.want {
			t.Errorf("Test failed. The input graph=%v is excepted to %t, but return %t", test.graph, test.want, got)
		}
	}
}
