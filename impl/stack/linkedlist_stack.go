package stack

import "fmt"

type stackNode[T any] struct {
	value T
	next  *stackNode[T]
}

type Stack[T any] struct {
	top  *stackNode[T] // points to the top node of the stack
	size int           // current size of the stack
}

func newStackNode[T any](val T) *stackNode[T] {
	return &stackNode[T]{
		value: val,
		next:  nil,
	}
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{top: nil, size: 0}
}

// Push the element into the stack
func (s *Stack[T]) Push(val T) {
	node := newStackNode(val)
	// insert at the first for O(1) time complexity operation
	node.next = s.top
	s.top = node
	s.size++
}

// Pop the element from the top
func (s *Stack[T]) Pop() (T, bool) {
	if s.top == nil {
		var zero T
		return zero, false
	}

	popped := s.top
	s.top = s.top.next
	s.size--
	return popped.value, true
}

// Peek the top element from the stack
func (s *Stack[T]) Peek() (T, bool) {
	if s.top == nil {
		var zero T
		return zero, false
	}

	return s.top.value, true
}

// Check if the stack is full or not
func (s *Stack[T]) Empty() bool {
	return s.top == nil
}

// Return the current size of the stack
func (s *Stack[T]) Size() int {
	return s.size
}

func (s *Stack[T]) String() string {
	out := ""
	curr := s.top

	for curr != nil {
		out += fmt.Sprintf("%v -> ", curr.value)
		curr = curr.next
	}

	return out
}
