# 1218. Longest Arithmetic Subsequence of Given Difference

Given an integer array `arr` and an integer `difference`, return the length of the longest subsequence in `arr` which is an arithmetic sequence such that the difference between adjacent elements in the subsequence equals `difference`.

A **subsequence** is a sequence that can be derived from `arr` by deleting some or no elements without changing the order of the remaining elements.

## Explain

本題思路是使用 `DP`, `dp[i]` 會紀錄當前最長的 `subsequence`
```
dp[i] = 1 + dp[i-diff] 
```

## Code

[Code](./solution.go)