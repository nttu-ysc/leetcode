package problem2289

func totalSteps(nums []int) int {
	var step int
	stack := []int{}
	for appendStack(nums, &stack); len(stack) > 0; appendStack(nums, &stack) {
		for len(stack) > 0 {
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nums = append(nums[:p], nums[p+1:]...)
		}
		step++
	}

	return step
}

func appendStack(nums []int, stack *[]int) {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			*stack = append(*stack, i)
		}
	}
}

func totalSteps2(nums []int) int {
	var step int

	for !nonDecreasingArr(nums) {
		step++
		tmp := nums[0]
		for i := 1; i < len(nums); {
			if tmp > nums[i] {
				tmp = nums[i]
				nums = append(nums[0:i], nums[i+1:]...)
				continue
			}
			tmp = nums[i]
			i++
		}
	}

	return step
}

func nonDecreasingArr(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}

	return true
}

func totalSteps3(nums []int) int {
	n := len(nums)
	DP := make([]int, n)
	Stack := make([]int, 1e5)
	top := -1
	ans := 0
	for i := n - 1; i >= 0; i-- {
		for top != -1 && nums[Stack[top]] < nums[i] {
			DP[i] = max(DP[Stack[top]], DP[i]+1)
			top--
		}
		ans = max(ans, DP[i])
		top++
		Stack[top] = i
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
