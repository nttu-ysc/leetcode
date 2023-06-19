package problem2739

import "testing"

func TestDistanceTraveled(t *testing.T) {
	tests := []struct {
		mainTank       int
		additionalTank int
		want           int
	}{
		{
			mainTank:       5,
			additionalTank: 10,
			want:           60,
		},
		{
			mainTank:       1,
			additionalTank: 2,
			want:           10,
		},
		{
			mainTank:       9,
			additionalTank: 1,
			want:           100,
		},
	}

	for _, test := range tests {
		if got := distanceTraveled(test.mainTank, test.additionalTank); got != test.want {
			t.Errorf("Test failed. The mainTank = %d, additionalTank = %d is expected to %d but return %d", test.mainTank, test.additionalTank, test.want, got)
		}
	}
}
