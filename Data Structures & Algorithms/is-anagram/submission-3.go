
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapS := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		mapS[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		fmt.Printf("%s: %d\n", string(t[i]), mapS[t[i]])
		if _, ok := mapS[t[i]]; !ok {
			return false
		}
		mapS[t[i]]--
		if mapS[t[i]] <= -1 {
			return false
		}
	}
	return true
}