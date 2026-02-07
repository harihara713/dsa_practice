package arrayandhashing

import (
	"reflect"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "Test - 1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "Test - 2",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "Test - 3",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := twoSum(tt.nums, tt.target)
			sort.Ints(tt.want)
			sort.Ints(res)

			if !reflect.DeepEqual(tt.want, res) {
				t.Fatalf("Expected = %v, Got = %v", tt.want, res)
			}
		})
	}
}
