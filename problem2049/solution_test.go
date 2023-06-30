package problem2049

import "testing"

func TestCountHighestScoreNodes(t *testing.T) {
	testCases := []struct {
		parents []int
		want    int
	}{
		{
			parents: []int{-1, 2, 0, 2, 0},
			want:    3,
		},
		{
			parents: []int{-1, 2, 0},
			want:    2,
		},
		{
			parents: []int{-1, 0},
			want:    2,
		},
		{
			parents: []int{-1, 3, 3, 5, 7, 6, 0, 0},
			want:    2,
		},
	}

	for _, tC := range testCases {
		if got := countHighestScoreNodes(tC.parents); got != tC.want {
			t.Errorf("Test failed. The input parents = %v is expected to %d but return %d", tC.parents, tC.want, got)
		}
	}
}
