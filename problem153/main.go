package problem153

func findMin(nums []int) int {
	l, h := 0, len(nums)-1
	target := nums[h]

	for l < h {
		mid := l + (h-l)/2
		if nums[mid] > target {
			l = mid + 1
		} else {
			h = mid
		}
	}
	return nums[h]
}
