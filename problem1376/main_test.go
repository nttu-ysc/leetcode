package problem1376

import "testing"

func TestNumOfMinutes(t *testing.T) {
	tests := []struct {
		n          int
		headID     int
		manager    []int
		informTime []int
		want       int
	}{
		{
			n:          1,
			headID:     0,
			manager:    []int{-1},
			informTime: []int{0},
			want:       0,
		},
		{
			n:          6,
			headID:     2,
			manager:    []int{2, 2, -1, 2, 2, 2},
			informTime: []int{0, 0, 1, 0, 0, 0},
			want:       1,
		},
	}

	for _, test := range tests {
		if got := numOfMinutes(test.n, test.headID, test.manager, test.informTime); got != test.want {
			t.Errorf("Test failed. The input n=%d headID=%d manager=%v informTime=%v is excepted to %d but return %d", test.n, test.headID, test.manager, test.informTime, test.want, got)
		}
	}

}
