package linkedlist

import (
	"testing"
)

func TestSinglyLinkedList_Append(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected []int
	}{
		{"append single", []int{10}, []int{10}},
		{"append multiple", []int{1, 2, 3}, []int{1, 2, 3}},
		{"append negative values", []int{-5, -10}, []int{-5, -10}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := NewSinglyLinkedList()
			for _, v := range tt.values {
				ll.Append(v)
			}

			if ll.Size() != len(tt.expected) {
				t.Fatalf("expected size %d, got %d", len(tt.expected), ll.Size())
			}

			// verify order
			for i, exp := range tt.expected {
				val, _ := ll.Get(i)
				if val != exp {
					t.Fatalf("at index %d: expected %d, got %d", i, exp, val)
				}
			}
		})
	}
}

func TestSinglyLinkedList_Insert(t *testing.T) {
	tests := []struct {
		name       string
		initial    []int
		pos        int
		value      int
		expectErr  bool
		finalState []int
	}{
		{
			"insert at head",
			[]int{2, 3},
			0,
			1,
			false,
			[]int{1, 2, 3},
		},
		{
			"insert at middle",
			[]int{1, 3},
			1,
			2,
			false,
			[]int{1, 2, 3},
		},
		{
			"insert at end",
			[]int{1, 2},
			2,
			3,
			false,
			[]int{1, 2, 3},
		},
		{
			"insert out of bounds",
			[]int{1, 2},
			5,
			10,
			true,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := NewSinglyLinkedList()
			for _, v := range tt.initial {
				ll.Append(v)
			}

			err := ll.Insert(tt.pos, tt.value)
			if tt.expectErr && err == nil {
				t.Fatalf("expected error but got nil")
			}

			if !tt.expectErr {
				if ll.Size() != len(tt.finalState) {
					t.Fatalf("expected size %d, got %d", len(tt.finalState), ll.Size())
				}

				for i, exp := range tt.finalState {
					got, _ := ll.Get(i)
					if got != exp {
						t.Fatalf("expected %d at pos %d, got %d", exp, i, got)
					}
				}
			}
		})
	}
}

func TestSinglyLinkedList_Get(t *testing.T) {
	ll := NewSinglyLinkedList()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	tests := []struct {
		pos       int
		expect    int
		expectErr bool
	}{
		{0, 1, false},
		{1, 2, false},
		{2, 3, false},
		{3, -1, true},
		{-1, -1, true},
	}

	for _, tt := range tests {
		got, err := ll.Get(tt.pos)
		if tt.expectErr {
			if err == nil {
				t.Errorf("expected error for pos %d but got nil", tt.pos)
			}
			continue
		}

		if got != tt.expect {
			t.Errorf("expected %d, got %d", tt.expect, got)
		}
	}
}

func TestSinglyLinkedList_Delete(t *testing.T) {
	tests := []struct {
		name       string
		initial    []int
		pos        int
		expectErr  bool
		finalState []int
	}{
		{
			"delete head",
			[]int{1, 2, 3},
			0,
			false,
			[]int{2, 3},
		},
		{
			"delete middle",
			[]int{1, 2, 3},
			1,
			false,
			[]int{1, 3},
		},
		{
			"delete last",
			[]int{1, 2, 3},
			2,
			false,
			[]int{1, 2},
		},
		{
			"delete out of bounds",
			[]int{1, 2},
			5,
			true,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := NewSinglyLinkedList()

			for _, v := range tt.initial {
				ll.Append(v)
			}

			err := ll.Delete(tt.pos)
			if tt.expectErr {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
				return
			}

			if ll.Size() != len(tt.finalState) {
				t.Fatalf("expected size %d, got %d", len(tt.finalState), ll.Size())
			}

			for i, exp := range tt.finalState {
				got, _ := ll.Get(i)
				if got != exp {
					t.Fatalf("expected %d at %d, got %d", exp, i, got)
				}
			}
		})
	}
}
