package problem207

import "testing"

func TestCanFinish(t *testing.T) {
	tests := []struct {
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{
			numCourses: 5,
			prerequisites: [][]int{
				{1, 0},
				{2, 1},
				{3, 1},
				{3, 2},
				{4, 3},
			},
			want: true,
		},
		{
			numCourses: 2,
			prerequisites: [][]int{
				{1, 0},
			},
			want: true,
		},
		{
			numCourses: 2,
			prerequisites: [][]int{
				{1, 0},
				{0, 1},
			},
			want: false,
		},
	}

	for _, test := range tests {
		if got := canFinish(test.numCourses, test.prerequisites); got != test.want {
			t.Errorf("Tets failed. The input numCourses=%d prerequisties = %v is expected to %t but return %t", test.numCourses, test.prerequisites, test.want, got)
		}
	}
}
