func twoSum(nums []int, target int) []int {
    seen := map[int]int{} //key = value, value = index

    for index, value := range nums {
        if value, ok := seen[target - value]; ok {
            return []int{value, index}
        }

        seen[value]=index
    }

    return nil
}
