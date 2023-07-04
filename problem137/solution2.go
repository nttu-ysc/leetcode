package problem137

func singleNumber2(nums []int) int {
	var ans int
	for i := 0; i < 32; i++ {
		sum := 0
		for _, num := range nums {
			sum += (num >> i) & 1
		}
		sum %= 3
		ans |= sum << i
	}
	if ans > 1<<31-1 {
		return ans - 1<<32
	}
	return ans
}
