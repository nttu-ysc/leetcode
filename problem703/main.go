package problem703

import "sort"

type KthLargest struct {
	kth  int
	nums []int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	return KthLargest{
		kth:  k,
		nums: nums,
	}
}

func (this *KthLargest) Add(val int) int {
	if len(this.nums) == 0 {
		this.nums = append(this.nums, val)
	}

	for i := 0; i < len(this.nums); i++ {
		if val > this.nums[i] {
			tmp := append([]int{}, this.nums[i:]...)
			this.nums = append(append(this.nums[:i], val), tmp...)
			break
		} else {
			if i == len(this.nums)-1 {
				this.nums = append(this.nums, val)
				break
			}
		}
	}

	return this.nums[this.kth-1]
}
