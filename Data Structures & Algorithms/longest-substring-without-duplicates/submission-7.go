func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	runes := []rune(s)
	left := 0
	right := 0
	m := make(map[rune]bool)
	maxCurr := 1
	
	for right < len(runes) {
		_, exists := m[runes[right]]
		if exists {
			maxCurr = max(maxCurr, right - left)
			for left < right {
				delete(m, runes[left])
				left++
				if runes[right] == runes[left - 1] {
					break
				}
			}
		}
		m[runes[right]] = true
		right++
	}
	return max(maxCurr, right - left)
}
