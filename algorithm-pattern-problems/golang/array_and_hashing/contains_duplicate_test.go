package arrayandhashing

import "testing"

func TestContainDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "Contains Duplicate - I",
			nums: []int{1, 2, 3, 1},
			want: true,
		},
		{
			name: "Contains Duplicate - II",
			nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			want: true,
		},
		{
			name: "Doesn't Contains Duplicate - I",
			nums: []int{1, 2, 3, 4},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := containsDuplicate(tt.nums)
			if res != tt.want {
				t.Fatalf("Expected = %t, got = %t", tt.want, res)
			}
		})
	}
}
