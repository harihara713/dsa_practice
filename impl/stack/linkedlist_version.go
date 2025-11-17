package main

import "fmt"

type stackNode struct {
	value interface{}
	next  *stackNode
}

type Stack struct {
	top  *stackNode // points to the top node of the stack
	size int        // current size of the stack
}

func newStackNode(val interface{}) *stackNode {
	return &stackNode{
		value: val,
		next:  nil,
	}
}

func NewStack() *Stack {
	return &Stack{top: nil, size: 0}
}

// Push the element into the stack
func (s *Stack) Push(val interface{}) {
	node := newStackNode(val)
	// insert at the first for O(1) time complexity operation
	node.next = s.top
	s.top = node
	s.size++
}

// Pop the element from the top
func (s *Stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}

	popped := s.top
	s.top = s.top.next
	s.size--
	return popped.value
}

// Peek the top element from the stack
func (s *Stack) Peek() interface{} {
	if s.top == nil {
		return nil
	}

	return s.top.value
}

// Check if the stack is full or not
func (s *Stack) Empty() bool {
	return s.top == nil
}

// Return the current size of the stack
func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) String() string {
	out := ""
	curr := s.top

	for curr != nil {
		out += fmt.Sprintf("%d -> ", curr.value)
		curr = curr.next
	}

	return out
}
