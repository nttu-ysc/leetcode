package problem347

func topKFrequent(nums []int, k int) []int {
	ans := make([]int, 0)
	dic := map[int]int{}
	for _, n := range nums {
		dic[n]++
	}

	redic := map[int][]int{}
	for n, count := range dic {
		redic[count] = append(redic[count], n)
	}

	for i := len(nums); len(ans) < k; i-- {
		for _, n := range redic[i] {
			ans = append(ans, n)
			if len(ans) == k {
				return ans
			}
		}
	}

	return ans
}
