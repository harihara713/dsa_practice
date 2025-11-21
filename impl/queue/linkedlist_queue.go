package queue

type queueNode[T any] struct {
	value T
	next  *queueNode[T]
}

func (n *queueNode[T]) Value() T {
	return n.value
}

func (n *queueNode[T]) Next() *queueNode[T] {
	return n.next
}

type Queue[T any] struct {
	front *queueNode[T]
	rear  *queueNode[T]
	size  int // current size of the queue
}

func newQueueNode[T any](val T) *queueNode[T] {
	return &queueNode[T]{value: val, next: nil}
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{front: nil, rear: nil, size: 0}
}

func (q *Queue[T]) Enqueue(value T) {
	node := newQueueNode[T](value)
	// if the queue is empty
	if q.front == nil {
		q.front = node
		q.rear = node
		q.size++
		return
	}

	q.rear.next = node
	q.rear = node
	q.size++
}

func (q *Queue[T]) Dequeue() T {
	if q.front == nil {
		var z T
		return z
	}

	d := q.front
	q.front = q.front.next
	q.size--
	return d.value
}

func (q *Queue[T]) Peek() T {
	if q.front == nil {
		var z T
		return z
	}

	return q.front.value
}

func (q *Queue[T]) Empty() bool {
	return q.front == nil
}

func (q *Queue[T]) Size() int {
	return q.size
}
