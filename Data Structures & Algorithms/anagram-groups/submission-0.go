
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, s := range strs {
		r := []rune(s)
		sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
		strSorted := string(r)
		m[strSorted] = append(m[strSorted], s)
	}
	anagrams := make([][]string, 0, len(m))
	for _, arr := range m {
		anagrams = append(anagrams, arr)
	}
	return anagrams
}
