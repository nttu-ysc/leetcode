# 2054. Two Best Non-Overlapping Events

You are given a **0-indexed** 2D integer array of events where $events[i] = [startTime_i, endTime_i, value_i]$. The $i^{th}$ event starts at $startTime_i$ and ends at $endTime_i$, and if you attend this event, you will receive a value of $value_i$. You can choose at most two non-overlapping events to attend such that the sum of their values is maximized.

Return this maximum sum.

Note that the start time and end time is inclusive: that is, you cannot attend two events where one of them starts and the other ends at the same time. More specifically, if you attend an event with end time `t`, the next event must start at or after `t + 1`.

---

## Explain

這題想求任兩個時段不重複的 `event` 相加的對大值  
整體思路如下:  
對 `event` 的 startTime 進行排序,  
逐一比較當前 `event` + 前 `i` 個 `event` 中的最大值  
但前 `i` 個 `event` 可能會重疊到,  
所以要用 `binary search` 來找最接近 `event` 的值

## Code

[Code](./solution.go)