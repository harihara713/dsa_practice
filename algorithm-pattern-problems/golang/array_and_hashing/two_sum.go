package arrayandhashing

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)

	for i, num := range nums {
		v, ok := m[num]
		if ok {
			return []int{i, v}
		}

		// add it to the map
		m[target-num] = i
	}

	return []int{-1, -1}
	// Time: O(n)
	// Space: O(n)
}
