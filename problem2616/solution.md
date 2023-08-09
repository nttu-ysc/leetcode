# 2616. Minimize the Maximum Difference of Pairs

You are given a 0-indexed integer array `nums` and an integer `p`. Find `p` pairs of indices of `nums` such that the maximum difference amongst all the pairs is minimized. Also, ensure no index appears more than once amongst the `p` pairs.

Note that for a pair of elements at the index `i` and `j`, the difference of this pair is `|nums[i] - nums[j]|`, where `|x|` represents the absolute value of `x`.

Return the ***minimum maximum** difference among all `p` pairs*. We define the maximum of an empty set to be zero.

## Explain

`nums` 可以形成 `p` 對 `pair`, 其中每對 `pair` 抓取的 `index` 不能相同  
題目要求要找到 `pair` 中相差最小的最大值

e.g. `nums = [10,1,2,7,1,3], p = 2`  
可以找到 max(|nums[1] - nums[4]|, |nums[2] - nums[5]|) = max(0, 1) = 1

該題思路可以使用 `binary search` 來查找  
對每個 `mid` 我們可以來輪詢 `nums` 看是否可以形成 `max 為 mid` 的 pair
```go
func canFormPairs(nums []int, mid int, p int) bool {
	var count int
	for i := 0; i < len(nums)-1 && count < p; {
		if nums[i+1]-nums[i] <= mid {
			count++
			i += 2
		} else {
			i++
		}
	}
	return count >= p
}
```

## Code

[Code](./solution.go)