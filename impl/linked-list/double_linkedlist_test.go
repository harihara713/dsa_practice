package main

import (
	"fmt"
	"reflect"
	"testing"
)

func listToSlice(l *DoublyLinkedList) []int {
	out := make([]int, 0, l.size)
	curr := l.head
	for curr != nil {
		out = append(out, curr.value)
		curr = curr.next
	}
	return out
}

func setupList(vals []int) *DoublyLinkedList {
	ll := NewDoublyLinkedList()
	for _, v := range vals {
		ll.Append(v)
	}
	return ll
}

func TestDoublyLinkedList_Append(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect []int
	}{
		{"append single", []int{10}, []int{10}},
		{"append multiple", []int{1, 2, 3}, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := NewDoublyLinkedList()

			for _, v := range tt.input {
				ll.Append(v)
			}

			got := listToSlice(ll)
			if !reflect.DeepEqual(got, tt.expect) {
				t.Fatalf("expected %v, got %v", tt.expect, got)
			}
		})
	}
}

func TestDoublyLinkedList_Insert(t *testing.T) {
	tests := []struct {
		name      string
		initial   []int
		pos       int
		val       int
		expectErr bool
		final     []int
	}{
		{"insert at head", []int{2, 3}, 0, 1, false, []int{1, 2, 3}},
		{"insert middle", []int{1, 3}, 1, 2, false, []int{1, 2, 3}},
		{"insert end", []int{1, 2}, 2, 3, false, []int{1, 2, 3}},
		{"insert out of bounds", []int{1, 2}, 5, 10, true, nil},
		{"insert on empty list at pos=0", []int{}, 0, 100, false, []int{100}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := setupList(tt.initial)

			err := ll.Insert(tt.pos, tt.val)
			if tt.expectErr {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
				return
			}

			got := listToSlice(ll)
			if !reflect.DeepEqual(got, tt.final) {
				t.Fatalf("expected %v, got %v", tt.final, got)
			}
		})
	}
}

func TestDoublyLinkedList_Get(t *testing.T) {
	tests := []struct {
		name      string
		initial   []int
		pos       int
		expectVal int
		expectErr bool
	}{
		{"get head", []int{10, 20, 30}, 0, 10, false},
		{"get middle", []int{10, 20, 30}, 1, 20, false},
		{"get last", []int{10, 20, 30}, 2, 30, false},
		{"get out of bounds", []int{1, 2}, 5, -1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := setupList(tt.initial)

			val, err := ll.Get(tt.pos)
			if tt.expectErr {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
				return
			}

			if val != tt.expectVal {
				t.Fatalf("expected %d, got %d", tt.expectVal, val)
			}
		})
	}
}

func TestDoublyLinkedList_Delete(t *testing.T) {
	tests := []struct {
		name      string
		initial   []int
		pos       int
		expectErr bool
		final     []int
	}{
		{"delete head", []int{1, 2, 3}, 0, false, []int{2, 3}},
		{"delete middle", []int{1, 2, 3}, 1, false, []int{1, 3}},
		{"delete last", []int{1, 2, 3}, 2, false, []int{1, 2}},
		{"delete only element", []int{5}, 0, false, []int{}},
		{"delete out of bounds", []int{1, 2}, 5, true, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ll := setupList(tt.initial)

			_, err := ll.Delete(tt.pos)
			if tt.expectErr {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
				return
			}

			got := listToSlice(ll)
			if !reflect.DeepEqual(got, tt.final) {
				t.Fatalf("expected %v, got %v", tt.final, got)
			}
		})
	}
}
