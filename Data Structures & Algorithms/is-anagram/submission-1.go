func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	runeS := []rune(s)
	runeT := []rune(t)
	mT := make(map[string]int, len(t))
	mS := make(map[string]int, len(s))
	count := 0
	for count < len(s) {
		charS := string(runeS[count])
		charT := string(runeT[count])
		mS[charS] += 1
		mT[charT] += 1
		count += 1
	}

	for i, _ := range mT {
		if mT[i] != mS[i] {
			return false
		}
	}
	return true
}
