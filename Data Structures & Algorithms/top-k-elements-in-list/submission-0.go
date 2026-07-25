func topKFrequent(nums []int, k int) []int {
	count := map[int]int{}

	for _, value := range nums {
		count[value]++
	}

	type pair struct{
		value int
		freq int
	}

	pairs := make([]pair, 0, len(count))
	for index, value := range count {
		pairs = append(pairs, pair{index, value})
	}

	sort.Slice(pairs, func(i, j int) bool {return pairs[i].freq > pairs[j].freq})

	result := []int{}
	for i := 0; i < k; i++ {
		result = append(result, pairs[i].value)
	}

	return result
}
