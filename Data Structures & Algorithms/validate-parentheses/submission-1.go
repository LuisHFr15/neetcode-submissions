func isValid(s string) bool {
    stack := []string{}
	corresp := map[string]string {
		")" : "(",
		"]": "[",
		"}": "{",
	}
	runes := []rune(s)

	for _, val := range runes {
		ch := string(val)
		if _, exists := corresp[ch]; !exists {
			stack = append(stack, ch)
		} else if len(stack) > 0 && stack[len(stack)-1] == corresp[ch] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}