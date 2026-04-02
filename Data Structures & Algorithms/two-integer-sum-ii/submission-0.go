func twoSum(nums []int, target int) []int {
	left := 0
    right := len(nums) - 1

    for left < right {
        sum := nums[left] + nums[right]
        if sum == target {
            return []int{left + 1, right + 1}
        }else if sum > target {
            right -= 1
        }else {
            left += 1
        }
    }

    return []int{-1, -1}
}
