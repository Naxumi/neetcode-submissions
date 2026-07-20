
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapS := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := mapS[s[i]]; ok {
			mapS[s[i]] = mapS[s[i]] + 1
			continue
		}
		mapS[s[i]] = 1
	}

	mapT := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		if _, ok := mapT[t[i]]; ok {
			mapT[t[i]] = mapT[t[i]] + 1
			continue
		}
		mapT[t[i]] = 1
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