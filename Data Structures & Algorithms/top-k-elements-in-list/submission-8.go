
func topKFrequent(nums []int, k int) []int {
	if len(nums) == 1 {
		return []int{nums[0]}
	}
	count := make(map[int]int)

	for _, num := range nums {
		count[num]++
	}
	
	freq := make([][]int, len(nums) + 1)

	for key, value := range count {
		freq[value] = append(freq[value], key)
	}

	var list []int
	for i := len(nums); i > 0; i-- {
		if len(list) == k {
			break
		}
		if len(freq[i]) == 0 {
			continue
		}
		for j := 0; j < len(freq[i]); j++ {
			list = append(list, freq[i][j])
		}
	}
	return list
}
