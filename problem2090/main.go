package problem2090

func getAverages(nums []int, k int) []int {
	ans := make([]int, len(nums))
	for r, sum := 0, 0; r < len(nums); r++ {
		sum += nums[r]
		if r-2*k > 0 {
			sum -= nums[r-2*k-1]
		}
		if r-k >= 0 {
			if r-2*k < 0 {
				ans[r-k] = -1
			} else {
				ans[r-k] = sum / (2*k + 1)
			}
		}
	}
	for i := 0; i < k && len(ans)-i > 0; i++ {
		ans[len(ans)-i-1] = -1
	}
	return ans
}

func getAverages2(nums []int, k int) []int {
	if k == 0 {
		return nums
	}
	div, sum := 2*k+1, 0

	prefix := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i < k || i >= len(nums)-k {
			prefix[i] = -1
			continue
		}
		if sum == 0 {
			for j := i - k; j <= i+k; j++ {
				sum += nums[j]
			}
		} else {
			sum -= nums[i-k-1]
			sum += nums[i+k]
		}
		prefix[i] = sum / div
	}
	return prefix
}
