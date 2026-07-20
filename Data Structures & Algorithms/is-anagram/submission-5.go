
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapS := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		mapS[s[i]]++
	}

	mapT := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		mapT[t[i]]++
	}

	for key, valS := range mapS {
		fmt.Printf("%s: %d\n", key, valS)
		if _, ok := mapT[key]; !ok {
			return false
		} else if valS != mapT[key] {
			return false
		}
	}
	return true
}