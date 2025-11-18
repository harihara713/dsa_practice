package stack

import "fmt"

type ArrayStack struct {
	data []int // store the data
	size int   // caoacity of the stack
	// pointing to the top of the stack
	top int
}

var (
	ErrStackOverflow = fmt.Errorf("stack is full")
	ErrStackEmpty    = fmt.Errorf("stack is empty")
)

func NewArrayStack(capacity int) *ArrayStack {
	return &ArrayStack{
		size: capacity,
		data: make([]int, capacity),
		top:  -1,
	}
}

// Push a value to the stack, if stack is full returns error
func (s *ArrayStack) Push(value int) error {
	if s.top+1 == s.size {
		return ErrStackOverflow
	}
	s.top++
	s.data[s.top] = value
	return nil
}

// Pop the top element from the stack
func (s *ArrayStack) Pop() (int, error) {
	if s.top == -1 {
		return -1, ErrStackEmpty
	}

	val := s.data[s.top]
	s.top--
	return val, nil
}

// Return the value at top of the stack, error if the stack is empty
func (s *ArrayStack) Peek() (int, error) {
	if s.top == -1 {
		return -1, ErrStackEmpty
	}

	return s.data[s.top], nil
}

// Check if the stack is empty or not
func (s *ArrayStack) Empty() bool {
	return s.top == -1
}

// Check if the stack is full or not
func (s *ArrayStack) Full() bool {
	return s.top+1 == s.size
}

// Return current size of the stack
func (s *ArrayStack) Size() int {
	return s.top + 1
}

// String representation of the stack
func (s *ArrayStack) String() string {
	if s.top == -1 {
		return ""
	}

	out := ""

	for i := s.top; i >= 0; i-- {
		if i == 0 {
			out += fmt.Sprintf("%d ", s.data[i])
			continue
		}
		out += fmt.Sprintf("%d -> ", s.data[i])
	}

	return out
}
