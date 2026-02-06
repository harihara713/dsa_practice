package arrayandhashing

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "Valid Anagram - I",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			name: "Not Valid Anagram - I",
			s:    "rat",
			t:    "car",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isAnagram(tt.s, tt.t)

			if res != tt.want {
				t.Fatalf("Expected = %t, got = %t", tt.want, res)
			}
		})
	}
}
