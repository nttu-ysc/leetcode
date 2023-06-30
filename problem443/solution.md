# 443. String Compression
## 題目

Given an array of characters chars, compress it using the following algorithm:

Begin with an empty string s. For each group of consecutive repeating characters in chars:

If the group's length is 1, append the character to s.
Otherwise, append the character followed by the group's length.
The compressed string s should not be returned separately, but instead, be stored in the input character array chars. Note that group lengths that are 10 or longer will be split into multiple characters in chars.

After you are done modifying the input array, return the new length of the array.

You must write an algorithm that uses only constant extra space.

---
## 方法

這個題目主要是對指標運用,整體的思路為  
紀錄相同char的初始位置`start`和char為多少`mem`, 並在壓縮後調整 *`start`*, *`mem`* 和 for-loop 的 *`i`* .

比較要注意的地方是, 我們是使用當前與前一個做比較來進行壓縮  
如果最後幾個char為連續字串的話, 我們會少壓縮一次  
因此需要在迴圈外再檢查一次來進行壓縮

[Code](./solution.go)