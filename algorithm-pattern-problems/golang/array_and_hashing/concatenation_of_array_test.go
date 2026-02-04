package arrayandhashing

import (
	"reflect"
	"testing"
)

func TestGetConcatenation(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "3 elements",
			nums: []int{1, 2, 1},
			want: []int{1, 2, 1, 1, 2, 1},
		},
		{
			name: "4 elements",
			nums: []int{1, 3, 2, 1},
			want: []int{1, 3, 2, 1, 1, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := getConcatenation(tt.nums)

			if len(ans) != len(tt.want) {
				t.Errorf("Want length = %d, got length = %d", len(tt.want), len(ans))
			}

			if !reflect.DeepEqual(tt.want, ans) {
				t.Errorf("Expected ans = %v, got ans = %v", tt.want, ans)
			}
		})
	}
}
