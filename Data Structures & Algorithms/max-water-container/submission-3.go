func maxArea(heights []int) int {
    left := 0
    right := len(heights) - 1
    ans := 0

    for left < right {
        area := (right - left) * Min(heights[left], heights[right])
        if area > ans {
            ans = area
        }

        if heights[left] < heights[right] {
            left++
        } else {
            right--
        }
    }

    return ans
}

func Min(a, b int) int {
    if a < b {
        return a
    }
    return b
}