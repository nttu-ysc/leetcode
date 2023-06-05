package problem300

import "sort"

func lengthOfLIS(nums []int) int {
	lis := []int{}

	for _, num := range nums {
		index := sort.SearchInts(lis, num)

		if index == len(lis) {
			lis = append(lis, num)
		} else if num < lis[index] {
			lis[index] = num
		}
	}

	return len(lis)
}
