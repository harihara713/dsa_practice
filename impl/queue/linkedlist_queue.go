package main

type queueNode struct {
	value interface{}
	next  *queueNode
}

type Queue struct {
	front *queueNode
	rear  *queueNode
	size  int // current size of the queue
}

func newQueueNode(val interface{}) *queueNode {
	return &queueNode{value: val, next: nil}
}

func NewQueue() *Queue {
	return &Queue{front: nil, rear: nil, size: 0}
}

func (q *Queue) Enqueue(value interface{}) {
	node := newQueueNode(value)
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

func (q *Queue) Dequeue() interface{} {
	if q.front == nil {
		return nil
	}

	d := q.front
	q.front = q.front.next
	q.size--
	return d.value
}

func (q *Queue) Peek() interface{} {
	if q.front == nil {
		return nil
	}

	return q.front.value
}

func (q *Queue) Empty() bool {
	return q.front == nil
}

func (q *Queue) Size() int {
	return q.size
}
