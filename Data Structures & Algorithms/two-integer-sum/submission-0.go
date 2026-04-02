func twoSum(nums []int, target int) []int {
    hMap := make(map[int]int)

    for i := 0; i < len(nums); i++ {
        diff := target - nums[i]
        if val, exists := hMap[diff]; exists {
            return []int{val, i}
        }else {
            hMap[nums[i]] = i
        }
    }
    return []int{-1,-1}
}
