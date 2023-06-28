package problem1514

import (
	"container/heap"
	"math"
)

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	visited := make([]bool, n)
	graph := make([]map[int]float64, n)
	for i := 0; i < len(graph); i++ {
		graph[i] = make(map[int]float64)
	}
	for k, e := range edges {
		graph[e[0]][e[1]] = -math.Log(succProb[k])
		graph[e[1]][e[0]] = -math.Log(succProb[k])
	}
	dist := make([]float64, n)
	for k := range dist {
		dist[k] = 1<<63 - 1
	}
	dist[start] = 0.0
	h := new(minHeap)

	heap.Push(h, el{
		dist: dist[start],
		node: start,
	})
	for h.Len() > 0 {
		cur := heap.Pop(h).(el)
		if visited[cur.node] {
			continue
		}
		visited[cur.node] = true
		if cur.node == end {
			return math.Exp(-dist[cur.node])
		}
		for next, prob := range graph[cur.node] {
			if visited[next] || dist[cur.node]+prob > dist[next] {
				continue
			}
			dist[next] = dist[cur.node] + prob
			heap.Push(h, el{
				dist: dist[next],
				node: next,
			})
		}
	}
	return 0
}

type el struct {
	dist float64
	node int
}

type minHeap []el

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].dist < h[j].dist
}

func (h *minHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *minHeap) Pop() interface{} {
	tmp := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return tmp
}

func (h *minHeap) Push(x interface{}) {
	(*h) = append((*h), x.(el))
}
