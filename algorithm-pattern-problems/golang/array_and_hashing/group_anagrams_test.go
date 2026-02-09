package arrayandhashing

import (
	"reflect"
	"slices"
	"strings"
	"testing"
)

func TestGroupAnagram(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "Test - I",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"},
			},
		},
		{
			name: "Test - II - Empty String Slice",
			strs: []string{""},
			want: [][]string{{""}},
		},
		{
			name: "Test - III - One string slice",
			strs: []string{"a"},
			want: [][]string{{"a"}},
		},
		{
			name: "Test - IV",
			strs: []string{"act", "pots", "tops", "cat", "stop", "hat"},
			want: [][]string{
				{"hat"}, {"act", "cat"}, {"stop", "pots", "tops"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := groupAnagrams(tt.strs)

			// Sort each inner slice
			for _, group := range tt.want {
				slices.Sort(group)
			}
			for _, group := range result {
				slices.Sort(group)
			}

			slices.SortFunc(tt.want, func(a, b []string) int {
				if len(a) == 0 || len(b) == 0 {
					return len(a) - len(b)
				}
				return strings.Compare(a[0], b[0])
			})
			slices.SortFunc(result, func(a, b []string) int {
				if len(a) == 0 || len(b) == 0 {
					return len(a) - len(b)
				}
				return strings.Compare(a[0], b[0])
			})

			if !reflect.DeepEqual(tt.want, result) {
				t.Errorf("Expected = %v, Got = %v\n", tt.want, result)
			}

		})
	}
}
