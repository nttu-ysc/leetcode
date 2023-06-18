package problem217

func containsDuplicate(nums []int) bool {
	hashMap := make(map[int]bool)
	for _, v := range nums {
		if _, ok := hashMap[v]; ok {
			return true
		}
		hashMap[v] = true
	}
	return false
}
