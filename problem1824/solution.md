# 1824. Minimum Sideway Jumps

There is a 3 lane road of length `n` that consists of `n + 1` points labeled from 0 to n. A frog starts at point 0 in the second lane and wants to jump to point n. However, there could be obstacles along the way.

You are given an array `obstacles` of length `n + 1` where each `obstacles[i]` (ranging from 0 to 3) describes an obstacle on the lane `obstacles[i]` at point `i`. If `obstacles[i] == 0`, there are no obstacles at point `i`. There will be at most one obstacle in the 3 lanes at each point.

- For example, if `obstacles[2] == 1`, then there is an obstacle on lane 1 at point 2.

The frog can only travel from point `i` to point `i + 1` on the same lane if there is not an obstacle on the lane at point `i + 1`. To avoid obstacles, the frog can also perform a side jump to jump to another lane (even if they are not adjacent) at the same point if there is no obstacle on the new lane.

- For example, the frog can jump from lane 3 at point 3 to lane 1 at point 3.

Return the *minimum number of side jumps* the frog needs to reach any lane at point n starting from lane `2` at point 0.

Note: There will be no obstacles on points `0` and `n`.

---
## Explain

我們需要長度為3的array `dp` 分別代表三個車道的最小花費  
最一開始青蛙位於車道2, 跳到隔壁車道花費為1, 所以 dp[1] = 0, dp[0] = ap[2] = 1  

如果當前位置 $i$ 沒有障礙物, 則 $i$ 和上一個位置 $i-1$ 花費相同  
如果當前位置有障礙物, 則下一個位置需判斷從隔壁跳過來的最小花費

`obj`為障礙物所在車道
```go
obj := obstacles[i]
pre0, pre1, pre2 := dp[0], dp[1], dp[2]
// 因為無法從障礙物車道跳到其他車道, 讓當前障礙物的車道花費為最大, 使其無法成為最小花費
for i := range dp {
    dp[i] = 1 << 32
}

// 如果障礙物不在當前車道, 則花費跟上一個位置相同
if obj != 1 {
    dp[0] = pre0
}
if obj != 2 {
    dp[1] = pre1
}
if obj != 3 {
    dp[2] = pre2
}

// 比較從隔壁跳過來的最小花費
if obj != 1 {
    dp[0] = min(dp[0], min(dp[1], dp[2])+1)
}
if obj != 2 {
    dp[1] = min(dp[1], min(dp[0], dp[2])+1)
}
if obj != 3 {
    dp[2] = min(dp[2], min(dp[0], dp[1])+1)
}
```

