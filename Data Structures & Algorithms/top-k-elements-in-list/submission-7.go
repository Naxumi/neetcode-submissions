
func topKFrequent(nums []int, k int) []int {
	if len(nums) == 1 {
		return []int{nums[0]}
	}
	count := make(map[int]int)

	for _, num := range nums {
		count[num]++
	}
	var sorted []int
	for num, _ := range count {
		sorted = append(sorted, num)
	}

	sort.Slice(sorted, func(i, j int) bool {
		return count[sorted[i]] > count[sorted[j]]
	})
	fmt.Println(sorted)

	var frequent []int
	for i := 0; i < k; i++ {
		frequent = append(frequent, sorted[i])
	}
	return frequent
}
