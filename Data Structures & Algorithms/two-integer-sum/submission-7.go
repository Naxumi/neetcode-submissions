func twoSum(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0, 1}
	}
    maps := make(map[int]int)

	for index, value := range nums {
		maps[value] = index
	}

	var i int
	var j int

	for index, value := range nums {
		valid := target - value
		if key, ok := maps[valid]; ok && key != index {
			fmt.Printf("%d: %d\n", key, )
			i = index
			j = maps[valid]
			break
		}
	}
	return []int{i, j}
}
