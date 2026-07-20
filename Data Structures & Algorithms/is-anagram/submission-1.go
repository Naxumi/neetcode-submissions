
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapS := make(map[string]int)
	for i := 0; i < len(s); i++ {
		if _, ok := mapS[string(s[i])]; ok {
			mapS[string(s[i])] = mapS[string(s[i])] + 1
			continue
		}
		mapS[string(s[i])] = 1
	}

	mapT := make(map[string]int)
	for i := 0; i < len(t); i++ {
		if _, ok := mapT[string(t[i])]; ok {
			mapT[string(t[i])] = mapT[string(t[i])] + 1
			continue
		}
		mapT[string(t[i])] = 1
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