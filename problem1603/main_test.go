package problem1603

import (
	"testing"
)

func TestParkingSystem(t *testing.T) {
	tests := []struct {
		init    []int
		actions []struct {
			addCar   int
			excepted bool
		}
	}{
		{init: []int{1, 1, 0},
			actions: []struct {
				addCar   int
				excepted bool
			}{
				{
					addCar:   1,
					excepted: true,
				},
				{
					addCar:   2,
					excepted: true,
				},
				{
					addCar:   3,
					excepted: false,
				},
				{
					addCar:   1,
					excepted: false,
				},
			},
		},
	}

	for _, test := range tests {
		parkingSystem := Constructor(test.init[0], test.init[1], test.init[2])
		for _, v := range test.actions {
			if res := parkingSystem.AddCar(v.addCar); res != v.excepted {
				t.Errorf("Test failed. The input %v is excepted to %t but return %t", v.addCar, v.excepted, res)
			}
		}
	}
}
