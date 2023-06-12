package problem228

import "testing"

func TestSummaryRanges(t *testing.T) {
	tests := []struct {
		nums []int
		want []string
	}{
		{
			nums: []int{0, 1, 2, 4, 5, 7},
			want: []string{"0->2", "4->5", "7"},
		},
		{
			nums: []int{0, 2, 3, 4, 6, 8, 9},
			want: []string{"0", "2->4", "6", "8->9"},
		},
		{
			nums: []int{},
			want: []string{},
		},
	}

	for _, test := range tests {
		got := summaryRanges(test.nums)

		if len(got) != len(test.want) {
			t.Errorf("Test failed.")
			break
		}

		for i := 0; i < len(got); i++ {
			if got[i] != test.want[i] {
				t.Errorf("Test failed. The nums=%v is excepted to %v but return %v", test.nums, test.want, got)
				break
			}
		}
	}
}
