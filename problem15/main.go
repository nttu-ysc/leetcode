package problem15

import "sort"

func threeSum(nums []int) [][]int {
	tripleMap := make(map[[3]int][]int)
	var ans = make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				triplet := [3]int{nums[i], nums[j], nums[k]}
				sort.Ints(triplet[:])

				if nums[i]+nums[j]+nums[k] == 0 {
					tripleMap[triplet] = []int{nums[i], nums[j], nums[k]}
				}
			}
		}
	}

	for _, v := range tripleMap {
		ans = append(ans, v)
	}

	return ans
}
