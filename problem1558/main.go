package problem1558

import "sort"

func minOperations(nums []int) int {
	var step int

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	// iterate until biggest number becomes zero
	for ; nums[0] > 0; step++ {
		for i := range nums {
			if nums[i]%2 == 0 {
				continue
			}
			// substract 1 from every odd number
			nums[i]--
			// count it as operation
			step++
		}

		// our biggest number is 0, exit loop
		if nums[0] == 0 {
			break
		}

		// divide all number in half
		for i := 0; i < len(nums); i++ {
			nums[i] /= 2
		}
	}
	return step
}
