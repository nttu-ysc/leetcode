# 859. Buddy Strings

Given two strings `s` and `goal`, return `true` if you can swap two letters in `s` so the result is equal to `goal`, otherwise, return `false`.

Swapping letters is defined as taking two indices `i` and `j` (0-indexed) such that `i != j` and swapping the characters at `s[i]` and `s[j]`.

For example, swapping at indices `0` and `2` in `"abcd"` results in `"cbad"`.

---

## Explain

如果 `s` 的長度 不等於 `goal` 的長度,直接返回 `false`  
- 找出 `s` 和 `goal` 不相等的字串位置存入 `diff` 陣列  
- 計算 `s` 裡面每個字元的數量存入 hashSet `hash` 中

---
1. 如果 `diff` 長度等於 `2` 判斷是否滿足
    > e.g.  
    > `s` = "ab", `goal` = "ba"

    ```go
    s[i] == goal[j] && s[j] == goal[i]
    ```

2. 如果 `diff` 長度等於 `0` 判斷 `s` 是否有重複的字元
    > e.g.  
    > `s` = "aa", `goal` = "aa"

    ```go
    len(hash) < len(s)
    ```

## 完整代碼
[Code](./solution.go)
