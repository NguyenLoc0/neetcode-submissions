func groupAnagrams(strs []string) [][]string {
    anagrams := map[string][]string{}
    
    for _, s := range strs {
        substring := []byte(s)
        sort.Slice(substring, func(i, j int) bool {return substring[i] < substring[j]})
        substring_sorted := string(substring)
        anagrams[substring_sorted] = append(anagrams[substring_sorted], s)
    }

    result := [][]string{}
    for _, anagram := range anagrams {
        result = append(result, anagram)
    }

    return result
}
