package problem1569

const MOD = 1000000007

type record struct {
	n int64
	k int64
}

func numOfWays(nums []int) int {
	dp := make(map[record]int64)
	return int((dfs(nums, dp) - 1) % MOD)
}

func dfs(nums []int, dp map[record]int64) int64 {
	if len(nums) <= 2 {
		return 1
	}
	l, r := []int{}, []int{}
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[0] {
			r = append(r, nums[i])
		} else if nums[i] < nums[0] {
			l = append(l, nums[i])
		}
	}

	c := combFunc(int64(len(l)+len(r)), int64(len(l)), dp)

	return c * dfs(l, dp) % MOD * dfs(r, dp) % MOD
}

func combFunc(n, k int64, dp map[record]int64) int64 {
	if k == 0 || n == k {
		return 1
	}

	record := record{
		n: n,
		k: k,
	}
	if dp[record] != 0 {
		return dp[record]
	}

	dp[record] = (combFunc(n-1, k-1, dp) + combFunc(n-1, k, dp)) % MOD
	return dp[record]
}
