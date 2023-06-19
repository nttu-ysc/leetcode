package problem560

func subarraySum(nums []int, k int) int {
	hashMap := make(map[int]int)
	hashMap[0] = 1
	var sum, res int
	for _, v := range nums {
		sum += v

		res += hashMap[sum-k]
		hashMap[sum]++
	}
	return res
}
