func isAnagram(s string, t string) bool {
	// If lengths don't match, they can't be anagrams
	if len(s) != len(t) {
		return false
	}

	// We only need one map! And we use 'byte' instead of 'string' for speed.
	counts := make(map[byte]int)

	// We can loop through both strings at the exact same time
	for i := 0; i < len(s); i++ {
		counts[s[i]]++ // Add 1 for the character in 's'
		counts[t[i]]-- // Subtract 1 for the character in 't'
	}

	// Verify: If they are anagrams, everything perfectly canceled out to 0
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}
