# 808. Soup Servings

There are two types of soup: type A and type B. Initially, we have `n` ml of each type of soup. There are four kinds of operations:
1. Serve `100` ml of soup A and `0` ml of soup B,
2. Serve `75` ml of soup A and `25` ml of soup B,
3. Serve `50` ml of soup A and `50` ml of soup B, and
4. Serve `25` ml of soup A and `75` ml of soup B.

When we serve some soup, we give it to someone, and we no longer have it. Each turn, we will choose from the four operations with an equal probability `0.25`. If the remaining volume of soup is not enough to complete the operation, we will serve as much as possible. We stop once we no longer have some quantity of both types of soup.

Note that we do not have an operation where all `100` ml's of soup B are used first.

Return the probability that soup A will be empty first, plus half the probability that A and B become empty at the same time. Answers within $10^{-5}$ of the actual answer will be accepted.

## Explain

該題目主要使用遞迴來查找答案, 可以先把四種運算方式定義出來
```golang
var operations = [][2]int{
	{100, 0},
	{75, 25},
	{50, 50},
	{25, 75},
}
```

接著開始計算每一種運算的機率, 比較需要注意的是這邊可以使用 `DP` 來加快遞迴的速度, 避免出現 `TLE` 的狀況

```golang
func helper(a, b int, dp *map[[2]int]float64) float64 {
	var p float64

	for _, operation := range operations {
		remainA, remainB := a-operation[0], b-operation[1]
		if remainA <= 0 && remainB > 0 {
			p += 0.25
			continue
		}

		if remainA <= 0 && remainB <= 0 {
			p += 0.25 * 0.5
			continue
		}

		if remainA > 0 && remainB > 0 {
			if pp, ok := (*dp)[[2]int{remainA, remainB}]; !ok {
				p += 0.25 * helper(remainA, remainB, dp)
			} else {
				p += 0.25 * pp
			}
			continue
		}
	}

	(*dp)[[2]int{a, b}] = p
	return p
}
```


## Code

[Code](./solution.go)
