package arrayandhashing

func containsDuplicate(nums []int) bool {
	if len(nums) == 1 {
		return false
	}

	hash := make(map[int]int, 0)

	for _, num := range nums {
		if _, ok := hash[num]; ok {
			return true
		}

		hash[num]++
	}

	return false
	// Time: O(n)
	// Space: O(n)
}
