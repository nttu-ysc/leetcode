# 1493. Longest Subarray of 1's After Deleting One Element

Given a binary array `nums`, you should delete one element from it.

Return the size of the longest *non-empty **subarray*** containing only `1`'s in the resulting array. Return `0` if there is no such subarray.

---

## Explain

依照題目, 我們要找出最長的子陣列, 最多可以包含 `1` 個 `0` .  
這題可以使用 `slide window` 來解

當我們的移動右邊界 `r` 直到子陣列超過 `1` 個 `0` 的時候,  
存入 `ans`
```
ans = max(ans, r-l-1)
```
移動左邊界 `l` 直到等於 `1` 個 `0` 的時候,  
繼續移動右邊界 `r`

## Code

[Code](./solution.go)