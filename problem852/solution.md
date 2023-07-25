# 852. Peak Index in a Mountain Array

An array arr a mountain if the following properties hold:

- `arr.length >= 3`
- There exists some `i` with `0 < i < arr.length - 1` such that:
    - `arr[0] < arr[1] < ... < arr[i - 1] < arr[i] `
    - `arr[i] > arr[i + 1] > ... > arr[arr.length - 1]`

Given a mountain array `arr`, return the index `i` such that `arr[0] < arr[1] < ... < arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1]`.

You must solve it in `O(log(arr.length))` time complexity.

## Explain

本題以 `arr` 表示一座山, 並且只有一個峰值  
如果直接遍歷 `arr` 的話也可以直接暴力解出來, 但是題目要求時間複雜度為

- `O(log(arr.length))`

很明顯可以看出是使用 `binary search` 來查找峰值

## Code

[Code](./solution.go)