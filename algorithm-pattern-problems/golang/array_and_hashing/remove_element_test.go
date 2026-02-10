package arrayandhashing

import (
	"reflect"
	"sort"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		wantK    int
		wantNums []int
	}{
		{
			name:     "Test-I",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			wantK:    2,
			wantNums: []int{2, 2, 3, 3},
		},
		{
			name:     "Test-II",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			wantK:    5,
			wantNums: []int{0, 1, 4, 0, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := removeElement(tt.nums, tt.val)

			if k != tt.wantK {
				t.Errorf("Expected K = %d, Got K = %d", tt.wantK, k)
			}

			sort.Slice(tt.nums[:k], func(i, j int) bool {
				return tt.nums[i] < tt.nums[j]
			})

			sort.Slice(tt.wantNums[:k], func(i, j int) bool {
				return tt.wantNums[i] < tt.wantNums[j]
			})

			if !reflect.DeepEqual(tt.nums[:k], tt.wantNums[:k]) {
				t.Fatalf("Expected nums = %v, Got nums = %v", tt.wantNums[:k], tt.nums[:k])
			}

		})
	}
}
