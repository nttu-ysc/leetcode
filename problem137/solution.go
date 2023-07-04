package problem137

func singleNumber(nums []int) int {
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hashMap[nums[i]]++
		if hashMap[nums[i]] == 3 {
			delete(hashMap, nums[i])
		}
	}
	for k := range hashMap {
		return k
	}
	return -1
}
