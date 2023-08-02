# 46. Permutations

Given an array `nums` of distinct integers, return all the *possible permutations*. You can return the answer in any order.

## Explain 

這題可以使用回朔法來解

對 `nums` 使用迴圈, 然後移除選中的 `num`, 把選中的 `num` 加到 `newTmp` 中
```go
for i := 0; i < len(nums); i++ {
    newNums := append([]int(nil), nums[:i]...)
    newNums = append(newNums, nums[i+1:]...)
    newTmp := append([]int(nil), tmp...)
    newTmp = append(newTmp, nums[i])
    helper(newNums, ans, newTmp)
}
```

遞迴結束條件為當前的 `nums` 長度為 `0` 就把 `tmp` 加入 `ans` 中

```go
if len(nums) == 0 {
    *ans = append(*ans, append([]int(nil), tmp...))
    return
}
```

## Code
[Code](./solution.go)