package linkedlist

import (
	"reflect"
	"testing"
)

//
// ---- Helpers ----
//

func setupCircularList(values []int) *CircularLinkedList {
	l := NewCircularLinkedList()
	for _, v := range values {
		l.Append(v)
	}
	return l
}

func circularToSlice(l *CircularLinkedList) []int {
	res := make([]int, 0, l.size)
	if l.head == nil {
		return res
	}

	curr := l.head
	for i := 0; i < l.size; i++ {
		res = append(res, curr.value)
		curr = curr.next
	}
	return res
}

//
// ---- Append Tests ----
//

func TestCircularLinkedList_Append(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		appendV  int
		expected []int
	}{
		{"append to empty", []int{}, 10, []int{10}},
		{"append single", []int{1}, 20, []int{1, 20}},
		{"append multiple", []int{1, 2}, 30, []int{1, 2, 30}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := setupCircularList(tt.initial)
			l.Append(tt.appendV)

			got := circularToSlice(l)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Fatalf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}

//
// ---- Insert Tests ----
//

func TestCircularLinkedList_Insert(t *testing.T) {
	tests := []struct {
		name      string
		initial   []int
		pos       int
		val       int
		expectErr bool
		expected  []int
	}{
		{"insert at head empty", []int{}, 0, 10, false, []int{10}},
		{"insert at head", []int{1, 2, 3}, 0, 10, false, []int{10, 1, 2, 3}},
		{"insert middle", []int{1, 2, 3}, 1, 10, false, []int{1, 10, 2, 3}},
		{"insert end", []int{1, 2, 3}, 3, 10, false, []int{1, 2, 3, 10}},
		{"insert out of bounds", []int{1, 2}, 5, 10, true, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := setupCircularList(tt.initial)

			err := l.Insert(tt.pos, tt.val)
			if tt.expectErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			got := circularToSlice(l)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Fatalf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}

//
// ---- Get Tests ----
//

func TestCircularLinkedList_Get(t *testing.T) {
	tests := []struct {
		name      string
		initial   []int
		pos       int
		val       int
		expectErr bool
	}{
		{"get head", []int{1, 2, 3}, 0, 1, false},
		{"get middle", []int{1, 2, 3}, 1, 2, false},
		{"get last", []int{1, 2, 3}, 2, 3, false},
		{"out of bounds", []int{1}, 5, -1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := setupCircularList(tt.initial)

			val, err := l.Get(tt.pos)
			if tt.expectErr {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
				return
			}

			if val != tt.val {
				t.Fatalf("expected %d, got %d", tt.val, val)
			}
		})
	}
}

//
// ---- Delete Tests ----
//

func TestCircularLinkedList_Delete(t *testing.T) {
	tests := []struct {
		name      string
		initial   []int
		pos       int
		expectErr bool
		expected  []int
	}{
		{"delete head", []int{1, 2, 3}, 0, false, []int{2, 3}},
		{"delete middle", []int{1, 2, 3}, 1, false, []int{1, 3}},
		{"delete last", []int{1, 2, 3}, 2, false, []int{1, 2}},
		{"delete only element", []int{5}, 0, false, []int{}},
		{"delete out of bounds", []int{1, 2}, 5, true, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := setupCircularList(tt.initial)

			_, err := l.Delete(tt.pos)
			if tt.expectErr {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			got := circularToSlice(l)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Fatalf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}
