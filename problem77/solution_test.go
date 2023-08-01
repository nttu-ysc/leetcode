package problem77

import "testing"

func TestCombine(t *testing.T) {
	testCases := []struct {
		n    int
		k    int
		want [][]int
	}{
		// {
		// 	n: 4,
		// 	k: 2,
		// 	want: [][]int{
		// 		{1, 2},
		// 		{1, 3},
		// 		{1, 4},
		// 		{2, 3},
		// 		{2, 4},
		// 		{3, 4},
		// 	},
		// },
		// {
		// 	n: 1,
		// 	k: 1,
		// 	want: [][]int{
		// 		{1},
		// 	},
		// },
		{
			n:    5,
			k:    4,
			want: [][]int{{1, 2, 3, 4}, {1, 2, 3, 5}, {1, 2, 4, 5}, {1, 3, 4, 5}, {2, 3, 4, 5}},
		},
	}

	for _, tC := range testCases {
		got := combine(tC.n, tC.k)
		if len(got) != len(tC.want) {
			t.Errorf("Test failed.")
			break
		}

		for k := range got {
			if len(got[k]) != len(tC.want[k]) {
				t.Errorf("Test failed.")
			}

			for kk := range got[k] {
				if got[k][kk] != tC.want[k][kk] {
					t.Errorf("Test failed. The input n = %d, k = %d is expected to %v but return %v", tC.n, tC.k, tC.want, got)
				}
				goto b
			}
		}
	b:
	}
}
