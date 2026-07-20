func hasDuplicate(nums []int) bool {
    isDuplicate := false
    for i := 0; i < len(nums); i++ {
        for j := 0; j < len(nums); j++ {
            if j == i {
                continue
            }
            if nums[i] == nums[j] {
                isDuplicate = true
                break
            }
        }
        if isDuplicate {
            break
        }
    }
    return isDuplicate
}
