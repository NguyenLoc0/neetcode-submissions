func twoSum(nums []int, target int) []int {
    find := map[int]int{}

	for index, value := range nums {
		if indexNeed, ok := find[target - value]; ok {
			return []int{indexNeed, index}
		}
		find[value] = index
	}
	return nil 
}
