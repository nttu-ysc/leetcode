package problem2462

import "testing"

func TestTotalCost(t *testing.T) {
	testCases := []struct {
		costs      []int
		k          int
		candidates int
		want       int64
	}{
		{
			costs:      []int{17, 12, 10, 2, 7, 2, 11, 20, 8},
			k:          3,
			candidates: 4,
			want:       11,
		},
		{
			costs:      []int{1, 2, 4, 1},
			k:          3,
			candidates: 3,
			want:       4,
		},
	}

	for _, tC := range testCases {
		if got := totalCost(tC.costs, tC.k, tC.candidates); got != tC.want {
			t.Errorf("Test failed. The input costs = %v, k = %d, candidates = %d, is expected to %d but return %d", tC.costs, tC.k, tC.candidates, tC.want, got)
		}
	}
}

var (
	costs      = []int{17, 12, 10, 2, 7, 2, 11, 20, 8}
	k          = 3
	candidates = 4
)

func BenchmarkSolution1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		totalCost(costs, k, candidates)
	}
}

func BenchmarkSolution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		totalCost2(costs, k, candidates)
	}
}
