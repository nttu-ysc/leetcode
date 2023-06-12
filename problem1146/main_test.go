package problem1146

import "testing"

func TestSnapshotArray(t *testing.T) {
	tests := []struct {
		actions []string
		inputs  [][]int
		want    []interface{}
	}{
		{
			actions: []string{"SnapshotArray", "set", "snap", "set", "get"},
			inputs: [][]int{
				{3},
				{0, 5},
				{},
				{0, 6},
				{0, 0},
			},
			want: []interface{}{nil, nil, 0, nil, 5},
		},
		{
			actions: []string{"SnapshotArray", "set", "snap", "snap", "snap", "get", "snap", "snap", "get"},
			inputs: [][]int{
				{1},
				{0, 15},
				{},
				{},
				{},
				{0, 2},
				{},
				{},
				{0, 0},
			},
			want: []interface{}{nil, nil, 0, 1, 2, 15, 3, 4, 15},
		},
	}

	for _, test := range tests {
		var arr SnapshotArray
		for k, action := range test.actions {
			switch action {
			case "SnapshotArray":
				arr = Constructor(test.inputs[k][0])
			case "set":
				arr.Set(test.inputs[k][0], test.inputs[k][1])
			case "snap":
				want := test.want[k].(int)
				if got := arr.Snap(); got != want {
					t.Errorf("Test failed. snap want index %d, but retrun %d", want, got)
				}
			case "get":
				want := test.want[k].(int)
				if got := arr.Get(test.inputs[k][0], test.inputs[k][1]); got != want {
					t.Errorf("Test failed. The index=%d snap_id=%d is excepted to %d but return %d", test.inputs[k][0], test.inputs[k][1], want, got)
				}
			}
		}
	}
}
