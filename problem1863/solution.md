# 1863. Sum of All Subset XOR Totals
## 題目
The XOR total of an array is defined as the bitwise `XOR` of all its elements, or `0` if the array is empty.
- For example, the XOR total of the array `[2,5,6]` is `2 XOR 5 XOR 6 = 1`.

Given an array `nums`, return the sum of all XOR totals for every subset of `nums`.  
Note: Subsets with the same elements should be counted multiple times.  
An array `a` is a subset of an array `b` if `a` can be obtained from `b` by deleting some (possibly zero) elements of `b`.

---

## 想法

題目要求所有子集的 `XOR` 的總和  
使用 `backtrack` 來遞迴可以求得所有的子集

每個 `nums` 的元素都會有種可能(包含和不包含)  
非別計算完之後相加就可以得到 `XOR` 總和  
遞迴結束條件直接返回 `XOR` 的直即可
```go
func backtrack(nums []int, index, num int) int {
	if index == len(nums) {
		return num
	}
	var ans int
	ans += backtrack(nums, index+1, num)
	ans += backtrack(nums, index+1, num^nums[index])
	return ans
}
```

## 完整代碼
[Code](./solution.go)
