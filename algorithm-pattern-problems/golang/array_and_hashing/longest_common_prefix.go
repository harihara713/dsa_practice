package arrayandhashing

import "sort"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	sort.Strings(strs)
	common := findCommon(strs[0], strs[len(strs)-1])

	return common
	// Time: O(nlogn)
	// Space: O(2*m)
}

func findCommon(s, t string) string {
	runeS := []rune(s)
	runeT := []rune(t)

	var common []rune
	i, j := 0, 0

	for i < len(runeS) && j < len(runeT) {
		if runeS[i] == runeT[j] {
			common = append(common, runeS[i])
		} else {
			break
		}
		i++
		j++
	}

	return string(common)
}
