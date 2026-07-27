func topKFrequent(nums []int, k int) []int {
	count := map[int]int {}

	for _, value := range nums{
		count[value]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, freq := range count{
		buckets[freq] = append(buckets[freq], num)
	}

	result := make([]int, 0, k)
	for freq := len(buckets)-1; freq >= 0 && k > 0; freq--{
		for _, value := range buckets[freq] {
			result = append(result, value)
			k--
		}
	}
	
	return result
}