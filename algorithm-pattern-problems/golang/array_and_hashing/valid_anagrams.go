package arrayandhashing

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// using two hash array or hash map
	// using hashmap
	m := make(map[rune]int, 0)

	for _, r := range s {
		m[r]++
	}

	for _, r := range t {
		m[r]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
	// Time: O(n)
	// space: O(n)
}
