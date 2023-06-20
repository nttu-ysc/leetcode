package problem42

import "testing"

func TestTrap(t *testing.T) {
	testCases := []struct {
		height []int
		want   int
	}{
		{
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			height: []int{4, 2, 0, 3, 2, 5},
			want:   9,
		},
	}
	for _, tC := range testCases {
		if got := trap(tC.height); got != tC.want {
			t.Errorf("Test failed. The height = %v is expected to %d but return %d", tC.height, tC.want, got)
		}
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	}
}
