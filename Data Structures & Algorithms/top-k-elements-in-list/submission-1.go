func topKFrequent(nums []int, k int) []int {
	count := map[int]int{}

	for _, value := range nums {
		count[value]++
	}

	bucket := make([][]int, len(nums)+1)
	for num, freq := range count {
		bucket[freq] = append(bucket[freq], num)
	}

	result := make([]int, 0, k)
	for i := len(bucket)-1; i>=0 && len(result) < k; i--{
		for _, value := range bucket[i] {
			result = append(result, value)
			if len(result) == k {
				break
			}
		}
	}
	return result
}
