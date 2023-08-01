# 77. Combinations

Given two integers `n` and `k`, return all possible combinations of `k` numbers chosen from the range `[1, n]`.

You may return the answer in any order.

## Explain

本題可以使用遞迴來解
```go
func helper(cur int, n int, k int, ans *[][]int, sel []int) {
	if len(sel) == k {
		*ans = append(*ans, append([]int{}, sel...))
		return
	}
	if cur > n {
		return
	}
	helper(cur+1, n, k, ans, sel)
	sel = append(sel, cur)
	helper(cur+1, n, k, ans, sel)
}
```

遞迴結束為 當 `arr` 長度到達 `k` 或者 `當前數字` 大於 `n` 就會題停止
```go
if len(sel) == k {
    *ans = append(*ans, append([]int{}, sel...))
    return
}
if cur > n {
    return
}
```

對每個數字會有兩種選擇 `取` 或 `不取`
```go
helper(cur+1, n, k, ans, sel)
sel = append(sel, cur)
helper(cur+1, n, k, ans, sel)
```

## Code
[Code](./solution.go)