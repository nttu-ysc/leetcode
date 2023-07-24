# 50. Pow(x, n)

Implement `pow(x, n)`, which calculates `x` raised to the power `n` (i.e., $x^n$).

## Explain

本題如果直接使用 `brute` 會有 `TLE` 的問題, 因此可以從指數 `n` 開始下手

當 `n < 0` , 結果為 $\frac{1}{x^{-n}}$  
因此我們這邊只要計算 `n > 0` 的情況就好

### e.g.
當 $n=10$  
所求
$$x^{10} = (x^{5})^{2} = (x * x^{4})^{2} = (x * (X^{2})^{2})^{2}  $$

從上述可以觀察出透過遞迴來查找最裡面的數字
```go
func helper(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	var res = helper(x, n>>1)
	res *= res
	if n%2 != 0 {
		res *= x
	}
	return res
}
```

## Code 

[Code](./solution.go)