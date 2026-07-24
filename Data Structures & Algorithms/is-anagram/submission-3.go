func isAnagram(s string, t string) bool {
    count := map[rune]int{}
    
    if len(s) != len(t) {
        return false
    }

    for _, char := range s {
        count[char]++
    }
    for _, char := range t {
        count[char]--
        if count[char]<0 {
            return false
        }
    }

    return true
}