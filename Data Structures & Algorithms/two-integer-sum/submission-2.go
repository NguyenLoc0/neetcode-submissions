func twoSum(nums []int, target int) []int {
    find := map[int]int{}

	for index, value := range nums {
		need := target - value
		if indexNeed, ok := find[need]; ok {
			return []int{indexNeed, index}
		}
		find[value] = index
	}
	return nil 
}
