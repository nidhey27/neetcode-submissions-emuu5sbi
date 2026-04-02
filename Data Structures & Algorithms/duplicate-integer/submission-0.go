func hasDuplicate(nums []int) bool {
    hMap := make(map[int]int)

    for i := 0; i < len(nums); i++ {
        if val, exists := hMap[nums[i]]; exists {
            hMap[nums[i]] = val + 1
        }else{
            hMap[nums[i]] = 1
        }
    }

    for _, val := range hMap {
        if val > 1 {
            return true
        }
    }

    return false
}
