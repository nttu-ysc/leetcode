# 2024. Maximize the Confusion of an Exam

A teacher is writing a test with n true/false questions, with 'T' denoting true and 'F' denoting false. He wants to confuse the students by maximizing the number of consecutive questions with the same answer (multiple trues or multiple falses in a row).

You are given a string answerKey, where answerKey[i] is the original answer to the ith question. In addition, you are given an integer k, the maximum number of times you may perform the following operation:

- Change the answer key for any question to 'T' or 'F' (i.e., set answerKey[i] to 'T' or 'F').

Return the maximum number of consecutive 'T's or 'F's in the answer key after performing the operation at most k times.

---

## Explain

`answerKey` 表示該考試的答案由 `T` 和 `F` 組成,  
我有有 `k` 次的機會可以更改答案, 使得找到最大連續相同答案

直覺使用 `slide window` 來查找, 會需要分成兩個case更改
- `T` 
- `F`

最後返回最大的長度

```go
// ch:
// -'T'
// -'F'
func helper(answerKey string, k int, ch byte) int {
	var ans, cnt int
	l, r := 0, 0
	for r < len(answerKey) {
		if answerKey[r] == ch {
			cnt++
		}
		for cnt == k+1 {
			ans = max(ans, r-l)
			if answerKey[l] == ch {
				cnt--
			}
			l++
		}
		r++
	}
	return max(ans, r-l)
}
```


## Code

[Code](./solution.go)