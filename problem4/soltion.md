# 4. Median of Two Sorted Arrays

Given two sorted arrays `nums1` and `nums2` of size `m` and `n` respectively, return the median of the two sorted arrays.

The overall run time complexity should be `O(log (m+n))`.

## Explain

本題是要找合併後中間項的值, 直接合併後再找中間項就可以了, 如果想提升速度的話, 可以在合併中就直接判斷是否已經有超過中間項的元素了, 如果有就直接返回

## Code

[Code](./solution.go)