package problem1396

import (
	"testing"
)

func TestUndergroundSystem(t *testing.T) {
	tests := []struct {
		actions  []string
		input    [][]interface{}
		excepted []interface{}
	}{
		{
			actions:  []string{"UndergroundSystem", "checkIn", "checkIn", "checkIn", "checkOut", "checkOut", "checkOut", "getAverageTime", "getAverageTime", "checkIn", "getAverageTime", "checkOut", "getAverageTime"},
			input:    [][]interface{}{{}, {45, "Leyton", 3}, {32, "Paradise", 8}, {27, "Leyton", 10}, {45, "Waterloo", 15}, {27, "Waterloo", 20}, {32, "Cambridge", 22}, {"Paradise", "Cambridge"}, {"Leyton", "Waterloo"}, {10, "Leyton", 24}, {"Leyton", "Waterloo"}, {10, "Waterloo", 38}, {"Leyton", "Waterloo"}},
			excepted: []interface{}{nil, nil, nil, nil, nil, nil, nil, 14.00000, 11.00000, nil, 11.00000, nil, 12.00000},
		},
		{
			actions: []string{"UndergroundSystem", "checkIn", "checkOut", "getAverageTime", "checkIn", "checkOut", "getAverageTime", "checkIn", "checkOut", "getAverageTime"},
			input: [][]interface{}{
				{}, {10, "Leyton", 3}, {10, "Paradise", 8}, {"Leyton", "Paradise"}, {5, "Leyton", 10}, {5, "Paradise", 16}, {"Leyton", "Paradise"}, {2, "Leyton", 21}, {2, "Paradise", 30}, {"Leyton", "Paradise"},
			},
			excepted: []interface{}{nil, nil, nil, 5.00000, nil, nil, 5.50000, nil, nil, 6.66667},
		},
		{
			actions: []string{"UndergroundSystem", "checkIn", "checkOut", "getAverageTime", "checkIn", "checkOut", "getAverageTime", "checkIn", "getAverageTime", "checkIn", "getAverageTime", "getAverageTime", "checkOut"},
			input: [][]interface{}{
				{}, {37043, "K2618O72", 29}, {37043, "U1DTINDT", 39}, {"K2618O72", "U1DTINDT"}, {779975, "K2618O72", 112}, {779975, "CN3K6CYM", 157}, {"K2618O72", "U1DTINDT"}, {706901, "K2618O72", 221}, {"K2618O72", "CN3K6CYM"}, {18036, "K2618O72", 258}, {"K2618O72", "U1DTINDT"}, {"K2618O72", "CN3K6CYM"}, {18036, "U1DTINDT", 293},
			},
			excepted: []interface{}{nil, nil, nil, 10.00000, nil, nil, 10.00000, nil, 45.00000, nil, 10.00000, 45.00000, nil},
		},
	}

	for _, test := range tests {
		us := Constructor()
		for i := 0; i < len(test.actions); i++ {
			switch test.actions[i] {
			case "UndergroundSystem":
			case "checkIn":
				us.CheckIn(test.input[i][0].(int), test.input[i][1].(string), test.input[i][2].(int))
			case "checkOut":
				us.CheckOut(test.input[i][0].(int), test.input[i][1].(string), test.input[i][2].(int))
			case "getAverageTime":
				at := us.GetAverageTime(test.input[i][0].(string), test.input[i][1].(string))
				if at != test.excepted[i].(float64) {
					t.Errorf("Test failed. average time is excepted to %f but response %f", test.excepted[i], at)
				}
			}
		}
	}
}
