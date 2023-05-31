package problem705

import (
	"testing"
)

func TestHashSet(t *testing.T) {
	tests := []struct {
		actions  []string
		inputs   []int
		excepted []interface{}
	}{
		{
			actions:  []string{"MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove", "contains"},
			inputs:   []int{0, 1, 2, 1, 3, 2, 2, 2, 2},
			excepted: []interface{}{nil, nil, nil, true, false, nil, true, nil, false},
		},
	}

	for _, test := range tests {
		var hs MyHashSet
		for i := 0; i < len(test.actions); i++ {
			switch test.actions[i] {
			case "MyHashSet":
				hs = Constructor()
			case "add":
				hs.Add(test.inputs[i])
			case "remove":
				hs.Remove(test.inputs[i])
			case "contains":
				if res := hs.Contains(test.inputs[i]); res != test.excepted[i] {
					t.Errorf("Test failed. contains is excepted to %t, but reture %t", test.excepted[i], res)
				}
			}
		}
	}
}
