package problem1603

import (
	"testing"
)

func TestParkingSystem(t *testing.T) {
	tests := []struct {
		init    []int
		actions []struct {
			addCar   int
			expected bool
		}
	}{
		{init: []int{1, 1, 0},
			actions: []struct {
				addCar   int
				expected bool
			}{
				{
					addCar:   1,
					expected: true,
				},
				{
					addCar:   2,
					expected: true,
				},
				{
					addCar:   3,
					expected: false,
				},
				{
					addCar:   1,
					expected: false,
				},
			},
		},
	}

	for _, test := range tests {
		parkingSystem := Constructor(test.init[0], test.init[1], test.init[2])
		for _, v := range test.actions {
			if res := parkingSystem.AddCar(v.addCar); res != v.expected {
				t.Errorf("Test failed. The input %v is expected to %t but return %t", v.addCar, v.expected, res)
			}
		}
	}
}
