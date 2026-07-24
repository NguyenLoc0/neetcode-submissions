func hasDuplicate(nums []int) bool {
    seen := map[int]int{}

    for _, value := range nums {
        seen[value]++;

        if seen[value] == 2 {
            return true
        }
    }

    return false
}
