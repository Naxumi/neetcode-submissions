func twoSum(nums []int, target int) []int {
    if len(nums) == 2 {
		return []int{0, 1}
	}

	maps := make(map[int]int)

	for key, value := range nums {
		difference := target - value
		if indexValue, ok := maps[difference]; ok {
			return []int{indexValue, key}
		}
		maps[value] = key
	}
	return []int{}
}
