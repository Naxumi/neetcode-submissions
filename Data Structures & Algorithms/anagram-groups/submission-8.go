func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{{strs[0]}}
	}

	lists := make(map[[26]int][]string)
	for _, str := range strs {
		var count [26]int
		for j := 0; j < len(str); j++ {
			count[str[j] - 'a']++
		}
		lists[count] = append(lists[count], str)
	}
	// fmt.Println(lists)
	sublists := [][]string{}
	for _, str := range lists {
		sublists = append(sublists, str)
	}

	return sublists
}
