# 209. Minimum Size Subarray Sum

Given an array of positive integers `nums` and a positive integer `target`, return the *minimal length* of a 
subarray
 *whose sum is greater than or equal to* `target`. If there is no such subarray, return `0` instead.

---

## Explain

本題的思路是用 `slide window` 來解, 透過對 `array` 一加一減來看最小可以達到 `target` 的長度

## Code
[Code](./solution.go)