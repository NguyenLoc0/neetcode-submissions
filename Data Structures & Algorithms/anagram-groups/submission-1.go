func groupAnagrams(strs []string) [][]string {
	anagrams := map[string][]string{}

	for _,s := range strs {
		b := []byte(s)
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		key := string(b)
		anagrams[key] = append(anagrams[key], s)
	}

	result := [][]string{}
	for _,anagram := range anagrams {
		result = append(result, anagram)
	}

	return result
}
