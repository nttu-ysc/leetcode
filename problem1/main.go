package problem1

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int, len(nums))
	for k, n := range nums {
		if index, existed := indexMap[target-n]; existed {
			return []int{index, k}
		}

		indexMap[n] = k
	}

	return []int{-1, -1}
}
