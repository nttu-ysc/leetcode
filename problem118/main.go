package problem118

func generate(numRows int) [][]int {
	ans := make([][]int, 0)
	for i := 1; i <= numRows; i++ {
		tmp := make([]int, i)
		for k := range tmp {
			if k == 0 || k == len(tmp)-1 {
				tmp[k] = 1
				continue
			}
			tmp[k] = ans[i-2][k] + ans[i-2][k-1]
		}
		ans = append(ans, tmp)
	}
	return ans
}
