package problem2462

import "container/heap"

func totalCost(costs []int, k int, candidates int) int64 {
	left := new(minHeap)
	right := new(minHeap)
	heap.Init(left)
	heap.Init(right)

	leftBoarder, rightBoarder := 0, len(costs)-1
	for i := 0; i < candidates && i < len(costs); i++ {
		heap.Push(left, costs[i])
		leftBoarder = i
	}

	for i := len(costs) - 1; i > len(costs)-1-candidates && i >= candidates; i-- {
		heap.Push(right, costs[i])
		rightBoarder = i
	}

	var total int64
	for k > 0 {
		k--
		if left.Len() < 1 {
			total += int64(heap.Pop(right).(int))
			continue
		}
		if right.Len() < 1 {
			total += int64(heap.Pop(left).(int))
			continue
		}
		l := heap.Pop(left).(int)
		r := heap.Pop(right).(int)
		if l <= r {
			total += int64(l)
			heap.Push(right, r)
			if leftBoarder+1 < rightBoarder {
				heap.Push(left, costs[leftBoarder+1])
				leftBoarder++
			}
		} else {
			total += int64(r)
			heap.Push(left, l)
			if rightBoarder-1 > leftBoarder {
				heap.Push(right, costs[rightBoarder-1])
				rightBoarder--
			}
		}
	}

	return total
}

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tmp
}
