func twoSum(nums []int, target int) []int {
    if len(nums) == 2 {
		return []int{0, 1}
	}

	var numsI int
	var numsJ int
	isFound := false
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				numsI = i
				numsJ = j
				isFound = true
				break
			}
		}
		if isFound {
			break
		}
	}
	return []int{numsI, numsJ}
}
