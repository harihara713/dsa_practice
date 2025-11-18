package dynamicarray

import "testing"

func TestAppend(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected []int
	}{
		{"append single", []int{10}, []int{10}},
		{"append multiple", []int{1, 2, 3}, []int{1, 2, 3}},
		{"append after resize", []int{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := NewDynamicArray()

			for _, v := range tt.values {
				arr.Append(v)
			}

			for i, expectedVal := range tt.expected {
				got, ok := arr.Get(i)
				if !ok {
					t.Fatalf("Get(%d) returned !ok", i)
				}
				if got != expectedVal {
					t.Errorf("expected %d at %d, got %d", expectedVal, i, got)
				}
			}

			if arr.Size() != len(tt.expected) {
				t.Errorf("expected size %d, got %d", len(tt.expected), arr.Size())
			}
		})
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name       string
		initial    []int
		pos        int
		insertVal  int
		want       []int
		shouldFail bool
	}{
		{
			name:      "insert middle",
			initial:   []int{1, 2, 3},
			pos:       1,
			insertVal: 99,
			want:      []int{1, 99, 2, 3},
		},
		{
			name:      "insert at end",
			initial:   []int{5, 6},
			pos:       2,
			insertVal: 10,
			want:      []int{5, 6, 10},
		},
		{
			name:       "insert out of bounds",
			initial:    []int{1, 2},
			pos:        5,
			insertVal:  50,
			shouldFail: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := NewDynamicArray()

			for _, v := range tt.initial {
				arr.Append(v)
			}

			err := arr.Insert(tt.pos, tt.insertVal)

			if tt.shouldFail {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
				return
			} else if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			for i, expectedVal := range tt.want {
				got, _ := arr.Get(i)
				if got != expectedVal {
					t.Errorf("expected %d at %d, got %d", expectedVal, i, got)
				}
			}
		})
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		name       string
		initial    []int
		pos        int
		newVal     int
		want       []int
		shouldFail bool
	}{
		{
			name:    "set middle",
			initial: []int{10, 20, 30},
			pos:     1,
			newVal:  99,
			want:    []int{10, 99, 30},
		},
		{
			name:       "set out of bounds",
			initial:    []int{1, 2},
			pos:        5,
			newVal:     10,
			shouldFail: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := NewDynamicArray()

			for _, v := range tt.initial {
				arr.Append(v)
			}

			err := arr.Set(tt.pos, tt.newVal)

			if tt.shouldFail {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
				return
			} else if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			for i, expectedVal := range tt.want {
				got, _ := arr.Get(i)
				if got != expectedVal {
					t.Errorf("expected %d at %d, got %d", expectedVal, i, got)
				}
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name       string
		initial    []int
		pos        int
		want       []int
		shouldFail bool
	}{
		{
			name:    "delete middle",
			initial: []int{1, 2, 3, 4},
			pos:     1,
			want:    []int{1, 3, 4},
		},
		{
			name:    "delete last",
			initial: []int{10, 20, 30},
			pos:     2,
			want:    []int{10, 20},
		},
		{
			name:       "delete out of bounds",
			initial:    []int{1, 2},
			pos:        10,
			shouldFail: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := NewDynamicArray()

			for _, v := range tt.initial {
				arr.Append(v)
			}

			err := arr.Delete(tt.pos)

			if tt.shouldFail {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
				return
			} else if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			for i, expectedVal := range tt.want {
				got, _ := arr.Get(i)
				if got != expectedVal {
					t.Errorf("expected %d at %d, got %d", expectedVal, i, got)
				}
			}

			if arr.Size() != len(tt.want) {
				t.Errorf("expected size %d, got %d", len(tt.want), arr.Size())
			}
		})
	}
}

func TestResize(t *testing.T) {
	arr := NewDynamicArray()

	for i := 0; i < 50; i++ {
		arr.Append(i)
	}

	if arr.Size() != 50 {
		t.Errorf("expected size 50, got %d", arr.Size())
	}

	val, ok := arr.Get(25)
	if !ok {
		t.Fatalf("Get(25) returned !ok")
	}

	if val != 25 {
		t.Errorf("expected 25 at index 25, got %d", val)
	}
}
