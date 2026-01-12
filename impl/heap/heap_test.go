package heap

import (
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name     string
		inserts  []int
		wantRoot int
		wantSize int
	}{
		{
			name:     "single element",
			inserts:  []int{42},
			wantRoot: 42,
			wantSize: 1,
		},
		{
			name:     "two elements - second larger",
			inserts:  []int{10, 25},
			wantRoot: 25,
			wantSize: 2,
		},
		{
			name:     "two elements - first larger",
			inserts:  []int{30, 15},
			wantRoot: 30,
			wantSize: 2,
		},
		{
			name:     "several elements with bubbling",
			inserts:  []int{10, 30, 20, 5, 60, 15, 45},
			wantRoot: 60,
			wantSize: 7,
		},
		{
			name:     "duplicates",
			inserts:  []int{7, 7, 7, 7},
			wantRoot: 7,
			wantSize: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New[int]()

			for _, val := range tt.inserts {
				h.Insert(val)
			}

			if h.size != tt.wantSize {
				t.Errorf("size = %d, want %d", h.size, tt.wantSize)
			}

			if len(h.data) == 1 || h.data[1] != tt.wantRoot {
				t.Errorf("root = %v, want = %d", h.data[1], tt.wantRoot)
			}

		})
	}
}

func TestDelete(t *testing.T) {
	inserts := []int{10, 30, 5, 25, 60, 15, 45, 1, 99, 70}

	h := New[int]()
	for _, v := range inserts {
		h.Insert(v)
	}

	// Expected descending order (max heap)
	expected := []int{99, 70, 60, 45, 30, 25, 15, 10, 5, 1}

	got := make([]int, 0, len(inserts))

	for h.size >= 1 { // while not empty
		max := h.Delete()
		got = append(got, max)
	}

	if len(got) != len(expected) {
		t.Fatalf("got %d elements, want %d", len(got), len(expected))
	}

	for i := range expected {
		if got[i] != expected[i] {
			t.Errorf("position %d: got %d, want %d", i, got[i], expected[i])
		}
	}
}

func TestDeleteOnEmptyHeap(t *testing.T) {
	h := New[int]()

	// Should panic or be handled — here we expect panic (as your code reads heaps[1])
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic on Delete() when empty, got none")
		}
	}()

	_ = h.Delete()
}

func TestMultipleDeleteAfterInserts(t *testing.T) {
	h := New[int]()
	values := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}

	for _, v := range values {
		h.Insert(v)
	}

	extracted := []int{}
	for h.size >= 1 {
		extracted = append(extracted, h.Delete())
	}

	// Just check it's sorted descending and contains same elements
	sortedOriginal := []int{9, 6, 5, 5, 4, 3, 2, 1, 1}
	if len(extracted) != len(sortedOriginal) {
		t.Fatalf("wrong count: got %d, want %d", len(extracted), len(sortedOriginal))
	}

	for i, want := range sortedOriginal {
		if extracted[i] != want {
			t.Errorf("extracted[%d] = %d, want %d", i, extracted[i], want)
		}
	}
}
