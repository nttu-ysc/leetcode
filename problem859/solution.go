package problem859

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	var diff []int
	hash := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		hash[s[i]]++
		if s[i] != goal[i] {
			diff = append(diff, i)
		}
	}
	if n := len(diff); n == 2 {
		return s[diff[0]] == goal[diff[1]] && s[diff[1]] == goal[diff[0]]
	} else if n == 0 {
		return len(hash) < len(s)
	}
	return false
}
