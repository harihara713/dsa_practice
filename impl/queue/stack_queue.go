package queue

import (
	"fmt"

	"github.com/harry713j/dsa_practice/impl/stack"
)

// Queue using stack
type StackQueue struct {
	first  *stack.Stack[int]
	second *stack.Stack[int]
}

func NewStackQueue() *StackQueue {
	return &StackQueue{
		first:  stack.NewStack[int](),
		second: stack.NewStack[int](),
	}
}

func (q *StackQueue) Enqueue(val int) {
	// push to the first stack
	q.first.Push(val)
}

func (q *StackQueue) Dequeue() (int, error) {
	// if the stacks are empty
	if q.first.Empty() && q.second.Empty() {
		return -1, fmt.Errorf("queue is empty")
	} else if q.second.Empty() {
		// transfer all the element to the second stack then pop
		for !q.first.Empty() {
			v, _ := q.first.Pop()
			q.second.Push(v)
		}

		val, _ := q.second.Pop()
		return val, nil
	} else {

		val, _ := q.second.Pop()
		return val, nil
	}
}

func (q *StackQueue) Peek() (int, error) {
	if q.first.Empty() && q.second.Empty() {
		return -1, fmt.Errorf("queue is empty")
	} else if q.second.Empty() {
		// transfer all the element to the second stack then pop
		for !q.first.Empty() {
			v, _ := q.first.Pop()
			q.second.Push(v)
		}

		val, _ := q.second.Peek()
		return val, nil
	} else {
		val, _ := q.second.Peek()
		return val, nil
	}
}

func (q *StackQueue) Empty() bool {
	return q.first.Empty() && q.second.Empty()
}

func (q *StackQueue) Size() int {
	return q.first.Size() + q.second.Size()
}
