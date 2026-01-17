package queue

import "fmt"

type ArrayQueue struct {
	data  []int // store the data
	size  int   // capacity of the queue
	front int   // pointing to front of the queue
	rear  int   // pointing to rear or end of the queue
}

func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{size: capacity, front: -1, rear: -1, data: make([]int, capacity)}
}

func (q *ArrayQueue) Enqueue(val int) error {
	if q.rear == q.size-1 {
		return fmt.Errorf("queue is full")
	}

	q.rear++
	q.data[q.rear] = val
	return nil
}

func (q *ArrayQueue) Dequeue() (int, error) {
	if q.rear == q.front {
		return -1, fmt.Errorf("queue is empty")
	}

	q.front++
	return q.data[q.front], nil
}

func (q *ArrayQueue) Peek() (int, error) {
	if q.rear == q.front {
		return -1, fmt.Errorf("queue is empty")
	}

	idx := q.front
	idx++
	return q.data[idx], nil
}

func (q *ArrayQueue) Empty() bool {
	return q.front == q.rear
}

func (q *ArrayQueue) Size() int {
	return q.rear + 1
}
