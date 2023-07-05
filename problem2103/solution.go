package problem2103

func countPoints(rings string) int {
	rods := make([]map[byte]bool, 10)
	for i := range rods {
		rods[i] = make(map[byte]bool)
	}
	for i := 0; i+1 < len(rings); i += 2 {
		rods[rings[i+1]-'0'][rings[i]] = true
	}
	var ans int
	for i := range rods {
		if len(rods[i]) == 3 {
			ans++
		}
	}
	return ans
}
