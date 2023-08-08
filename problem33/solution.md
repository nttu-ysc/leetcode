# 33. Search in Rotated Sorted Array

There is an integer array `nums` sorted in ascending order (with distinct values).

Prior to being passed to your function, `nums` is possibly rotated at an unknown pivot index `k` (`1 <= k < nums.length`) such that the resulting array is `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]` (0-indexed). For example, `[0,1,2,4,5,6,7]` might be rotated at pivot index `3` and become `[4,5,6,7,0,1,2]`.

Given the array nums after the possible rotation and an integer `target`, return the index of `target` if it is in `nums`, or `-1` if it is not in nums.

You must write an algorithm with `O(log n)` runtime complexity.

## Explain

本題目給定一個 array `nums` 是依照 `asc` 進行排序的, 再把 array 翻轉 `n` 次 (n 為未知的數字)  
給定一個 `target`, 如果 `nums` 中有 `target` 返回 `index` 否則返回 `-1`

題目要求時間複雜度為 `O(log n)` 可以猜出應該是用 `binary search`  
因為 `nums` 有翻轉過, 所以使用 `binary search` 之後, 要在判斷 `corner value` 是否在範圍內

## Code

[Code](./solution.go)