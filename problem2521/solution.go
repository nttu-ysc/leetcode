package problem2521

func distinctPrimeFactors(nums []int) int {
	hashMap := make(map[int]int)
	for i := range nums {
		h := helper(nums[i])
		for k := range h {
			hashMap[k] += h[k]
		}
	}
	return len(hashMap)
}

func helper(num int) map[int]int {
	hashMap := make(map[int]int)
	for i := 2; num > 1; i++ {
		for num%i == 0 {
			hashMap[i]++
			num /= i
		}
	}
	return hashMap
}
