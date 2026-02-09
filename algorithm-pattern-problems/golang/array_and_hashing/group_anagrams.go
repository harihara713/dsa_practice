package arrayandhashing

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}

	for _, s := range strs {
		rS := []rune(s)
		sort.Slice(rS, func(i, j int) bool {
			return rS[i] < rS[j]
		})

		sortedS := string(rS)
		m[sortedS] = append(m[sortedS], s)
	}

	ans := [][]string{}
	for _, v := range m {
		ans = append(ans, v)
	}

	return ans
	// Time: O(n * m log m)
	// Space: O(n)
}
