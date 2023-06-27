package problem373

import (
	"container/heap"
)

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var ans [][]int
	h := new(minHeap)
	heap.Init(h)
	for k := range nums1 {
		heap.Push(h, [3]int{nums1[k] + nums2[0], k, 0})
	}
	for len(ans) < min(k, len(nums1)*len(nums2)) {
		tmp := heap.Pop(h).([3]int)
		i, j := tmp[1], tmp[2]
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < len(nums2) {
			heap.Push(h, [3]int{nums1[i] + nums2[j+1], i, j + 1})
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type minHeap [][3]int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.([3]int))
}

func (h *minHeap) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tmp
}
