func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{{strs[0]}}
	}

	isAnagram := make(map[int]bool)
	var sublists [][]string
	for i := 0; i < len(strs); i++ {
		// if isAnagram[i + 1] {
		// 	fmt.Printf("isAnagram: %d | %d\n", i, i + 1)
		// 	continue
		// }
        // fmt.Printf("%d: %s | %d: %s\n", i, strs[i], i + 1, strs[i + 1])
		if isAnagram[i] {
			continue
		}

		list := []string{}
		for j := i + 1; j < len(strs); j++ {
			if len(strs[i]) != len(strs[j]) {
				continue
			}
			count := [26]int{}
			for k := 0; k < len(strs[j]); k++ {
				count[strs[i][k] - 'a']++
				count[strs[j][k] - 'a']--
			}
			isValidAnagram := true
			// fmt.Printf("i: %s | j: %s - count : %v\n", strs[i], strs[j], count)
			for _, value := range count {
				if value != 0 {
					isValidAnagram = false
					break
				}
			}
			if !isValidAnagram {
				continue
			}
			// sublists = append(sublists, []string{strs[j]})
			isAnagram[i] = true
			isAnagram[j] = true
			list = append(list, strs[j])
			// fmt.Printf("i: %s | j: %s\n", strs[i], strs[j])
		}
		if !isAnagram[i] {
			sublists = append(sublists, []string{strs[i]})
		} else {
			list = append(list, strs[i])
		}
		if len(list) != 0 {
		sublists = append(sublists, list)
		}

		// for j := 0; j < len(strs[i]); j++ {
		// 	count[strs[i][j] - 'a']++
		// 	count[strs[i+1][j] - 'a']--
		// }

		// isValidAnagram := true
		// for _, v := range strs[i] {
		// 	if count[v - 'a'] != 0 {
		// 		isValidAnagram = false
		// 		break
		// 	}
		// }
		// if !isValidAnagram {
		// 	fmt.Printf("discontinued: %d\n", i)
		// 	continue
		// }
		// isAnagram[i] = true
		// isAnagram[i+1] = true
		// sublists[i] = append(sublists[i], strs[i + 1])
		// fmt.Println("hii")
		// fmt.Println(sublists[i])
	}

	return sublists
}
