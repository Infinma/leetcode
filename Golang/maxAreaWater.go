func maxArea(height []int) int {
	var area int
	var min int
	i := 0
	j := len(height) - 1
	sol := 0
	for i < j {
			min = height[i]
			if height[j] < height[i] {
					min = height[j]
			}
			area = min * (j - i)
			if (sol < area) {
					sol = area
			}
			if height[i] < height[j] {
					i += 1
			} else {
					j -= 1
			}
	}
	return sol
}
