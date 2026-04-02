func maxArea(heights []int) int {
	l := 0
	r := len(heights) - 1
	max := 0
	for l < r {
		h := MinInt(heights[l], heights[r])
		max = MaxInt(max, ((r - l) * h))

		if heights[l] < heights[r] {
			l += 1
		} else {
			r -= 1
		}
	}
	return max
}

func MinInt(a,b int) int {
	if a > b {
		return b
	}
	return a
}

func MaxInt(a,b int) int {
	if a > b {
		return a
	}
	return b
}


