package problem1187

import (
	"sort"
)

const INF = 1 << 30

type node struct {
	index int
	prev  int
}

var dp map[node]int

func makeArrayIncreasing(arr1, arr2 []int) int {
	sort.Ints(arr2)
	for i := 1; i < len(arr2); {
		if arr2[i] == arr2[i-1] {
			arr2 = append(arr2[:i], arr2[i+1:]...)
			continue
		}
		i++
	}

	dp = make(map[node]int)
	res := solve(arr1, arr2, 0, -1)

	if res == INF {
		return -1
	}

	return res
}

func solve(arr1, arr2 []int, i, prev int) int {
	if i >= len(arr1) {
		return 0
	}

	if v, ok := dp[node{i, prev}]; ok {
		return v
	}

	res1, res2 := INF, INF
	newCurr := binarySearch(arr2, prev)
	if newCurr > prev {
		res1 = solve(arr1, arr2, i+1, newCurr) + 1
		res1 = min(res1, INF)
	}

	if arr1[i] > prev {
		res2 = solve(arr1, arr2, i+1, arr1[i])
	}

	dp[node{i, prev}] = min(res1, res2)
	return dp[node{i, prev}]
}

func binarySearch(arr []int, target int) int {
	res := -1
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] <= target {
			l = mid + 1
		} else {
			res = arr[mid]
			r = mid - 1
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
