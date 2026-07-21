func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var amount [26]int

	for i := 0; i < len(s); i++ {
		amount[s[i] - 'a']++
		amount[t[i] - 'a']--
	}

	for _, val := range amount {
		if val != 0 {
			return false
		}
	}
	return true
}
