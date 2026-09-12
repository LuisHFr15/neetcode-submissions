func maxArea(heights []int) int {
	var maxVolume int

	left := 0
	right := len(heights) - 1

	for left < right {
		vol := min(heights[left], heights[right]) * (right - left)
		if vol >= maxVolume {
			maxVolume = vol
		}
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return maxVolume
}
