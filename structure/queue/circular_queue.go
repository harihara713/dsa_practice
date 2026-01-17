package queue

import "fmt"

type CircularQueue struct {
	data  []int
	size  int // capacity of the queue
	front int
	rear  int
}

func NewCircularQueue(capacity uint) *CircularQueue {
	return &CircularQueue{size: int(capacity), front: 0, rear: 0, data: make([]int, capacity)}
}

func (q *CircularQueue) Enqueue(value int) error {
	// if the queue is full
	if (q.rear+1)%q.size == q.front {
		return fmt.Errorf("queue is full")
	}

	q.rear = (q.rear + 1) % q.size // for proper circular behavior
	q.data[q.rear] = value
	return nil
}

func (q *CircularQueue) Dequeue() (int, error) {
	if q.front == q.rear {
		return -1, fmt.Errorf("queue is empty")
	}

	q.front = (q.front + 1) % q.size
	return q.data[q.front], nil
}

func (q *CircularQueue) Peek() (int, error) {
	if q.front == q.rear {
		return -1, fmt.Errorf("queue is empty")
	}

	idx := q.front
	idx = (idx + 1) % q.size
	return q.data[idx], nil
}

func (q *CircularQueue) Empty() bool {
	return q.front == q.rear
}

func (q *CircularQueue) Size() int {
	if q.rear == q.front {
		return 0
	} else if q.rear > q.front {
		return q.rear - q.front
	} else {
		return q.size - q.front + q.rear
	}
}
