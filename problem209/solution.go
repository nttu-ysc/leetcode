package problem209

func minSubArrayLen(target int, nums []int) int {
	var count int = 10e5
	var sum int
	var sub bool
	for l, r := 0, 0; r < len(nums); r++ {
		sum += nums[r]
		for sum >= target {
			count = min(count, r-l+1)
			sum -= nums[l]
			l++
			sub = true
		}
	}
	if !sub {
		return 0
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
