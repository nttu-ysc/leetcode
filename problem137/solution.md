# 137. Single Number II

Given an integer array `nums` where every element appears three times except for one, which appears exactly once. Find the single element and return it.

You must implement a solution with a linear runtime complexity and use only constant extra space.

**Constraints:**
- $1$ <= nums.length <= $3 * 10^4$
- $-2^{31}$ <= nums[i] <= $2^{31} - 1$
- Each element in `nums` appears exactly **three times** except for one element which appears **once**.

---

## Approach 1

我們先用 `hashMap` 來存每個元素的個數, 並直接判斷當個數 = 3 的時候刪掉該 key  
最後 `hashMap` 會只剩下個數為 1 的元素, 即為所求

## Code
[Code](./solution.go)

---

## Approach 2

我可以對每個 `num` 的 bit 總和 `sum` 對 `3` 進行 `mod`, 取完 `mod` 之後 剩下的就是 single num 的 bit 值, 再與 `ans` 取 `|` 運算

e.g. `nums = [1, 2, 1, 1]`,  
`i = 0`,  
```
sum = (1 >> 0 & 1) + (2 >> 0 & 1) + (1 >> 0 & 1) + (1 >> 0 & 1)
sum = 1 + 0 + 1 + 1
sum = 3

// sum % 3
sum = 0
ans = 0 | 0 << 0
ans = 0
```
可以取得 ans 為 0

`i = 1`,  
```
sum = (1 >> 1 & 1) + (2 >> 1 & 1) + (1 >> 1 & 1) + (1 >> 1 & 1)
sum = 0 + 1 + 0 + 0
sum = 1

// sum % 3
sum = 1
ans = 0 | 1 << 1
// 二進位制表示為 10, 十進位為 2
ans = 2
```
可以取得 ans 為 2

## Code
[Code](./solution2.go)
