package arrayandhashing

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{
			name: "Test-I",
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "Test-II",
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			name: "Test-III",
			strs: []string{"fight", "flight", "flower", "flow"},
			want: "f",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := longestCommonPrefix(tt.strs)

			if res != tt.want {
				t.Errorf("Expected Common Prefix = %s, Got = %s", tt.want, res)
			}
		})
	}

}
