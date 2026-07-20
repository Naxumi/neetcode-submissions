func hasDuplicate(nums []int) bool {
    sets := make(map[int]bool)
    for i := 0; i < len(nums); i++ {
        if sets[nums[i]] {
            return true
        }
        sets[nums[i]] = true
    }
    return false
}
