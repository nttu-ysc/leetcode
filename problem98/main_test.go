package problem98

import (
	"testing"
)

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		root []interface{}
		want bool
	}{
		{
			root: []interface{}{2, 1, 3},
			want: true,
		},
		{
			root: []interface{}{5, 1, 4, nil, nil, 3, 6},
			want: false,
		},
		{
			root: []interface{}{2, 2, 2},
			want: false,
		},
		{
			root: []interface{}{5, 4, 6, nil, nil, 3, 7},
			want: false,
		},
	}

	for _, test := range tests {
		treeNode := construct(test.root)
		if got := isValidBST(treeNode); got != test.want {
			t.Errorf("Test failed. the input root=%v is excepted to %t but return %t", test.root, test.want, got)
		}
	}

}
