package problem2542

import (
	"container/heap"
	"sort"
)

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	var sum int64
	var ans int64
	var arr [][]int
	for i := 0; i < len(nums1); i++ {
		arr = append(arr, []int{nums1[i], nums2[i]})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})

	minHeap := &MinHeap{}

	for i := 0; i < k; i++ {
		heap.Push(minHeap, arr[i])
		sum += int64(arr[i][0])
	}

	ans = sum * int64(arr[k-1][1])

	for i := k; i < len(arr); i++ {
		smallNum := heap.Pop(minHeap).([]int)
		sum += int64(arr[i][0] - smallNum[0])

		heap.Push(minHeap, arr[i])

		ans = max(ans, sum*int64(arr[i][1]))
	}

	return ans
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}

	return b
}

type MinHeap [][]int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
