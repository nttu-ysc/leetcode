# 17. Letter Combinations of a Phone Number

Given a string containing digits from `2-9` inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

## Explain

這題要得蠻單純的, 直接暴力解也不會花太久的時間
暴力解就直接把按鍵上的每個字母組合相加

另外還可以使用回朔法來迭代, 當組合結果 `tmp` 的字串長度 和 `digits` 的字串長度相等, 就把 `tmp` 加入到 `ans` 中, 並且遍歷取出的字符串即為所求


## Code

[approach 1 brute-force](solution.go)

[approach 2: backtrack](solution2.go)