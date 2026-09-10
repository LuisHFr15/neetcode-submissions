func isPalindrome(s string) bool {
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	str := re.ReplaceAllString(s, "")
	runes := []rune(strings.ToLower(str))
	for in := 0; in < len(runes) / 2; in++ {
		secIn := len(runes) - 1 - in
		if runes[in] != runes[secIn] {
			fmt.Println(in)
			fmt.Println(runes[in])
			fmt.Println(runes[secIn])
			return false
		}
	}

	return true
}
