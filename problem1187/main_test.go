package problem1187

import "testing"

func TestMakeArrayIncreasing(t *testing.T) {
	tests := []struct {
		arr1 []int
		arr2 []int
		want int
	}{
		{
			arr1: []int{1, 5, 3, 6, 7},
			arr2: []int{1, 3, 2, 4, 2, 2, 2},
			want: 1,
		},
		{
			arr1: []int{1, 5, 3, 6, 7},
			arr2: []int{4, 3, 1},
			want: 2,
		},
		{
			arr1: []int{1, 5, 3, 6, 7},
			arr2: []int{1, 6, 3, 3},
			want: -1,
		},
		{
			arr1: []int{9, 3},
			arr2: []int{5, 6},
			want: 2,
		},
	}

	for _, test := range tests {
		var arr1, arr2 []int
		arr1 = append(arr1, test.arr1...)
		arr2 = append(arr2, test.arr2...)
		if got := makeArrayIncreasing(test.arr1, test.arr2); got != test.want {
			t.Errorf("Test failed. The input arr1=%v arr2=%v is excepted to %d but return %d", arr1, arr2, test.want, got)
		}
	}
}
