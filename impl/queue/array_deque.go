package main

import "fmt"

type ArrayDeque struct {
	data  []int
	size  int // capacity
	front int
	rear  int
}

func NewArrayDeque(capacity uint) *ArrayDeque {
	return &ArrayDeque{
		size:  int(capacity),
		front: -1,
		rear:  -1,
		data:  make([]int, capacity),
	}
}

func (d *ArrayDeque) EnqueueFirst(val int) error {
	if d.front == -1 {
		return fmt.Errorf("deque is full from front-end")
	}

	d.data[d.front] = val
	d.front--
	return nil
}

func (d *ArrayDeque) EnqueueLast(val int) error {
	// when the deque is full
	if d.rear == d.size-1 {
		return fmt.Errorf("deque is full from rear-end")
	}

	d.rear++
	d.data[d.rear] = val
	return nil
}

func (d *ArrayDeque) DequeueFirst() (int, error) {
	// empty deque
	if d.front == d.rear {
		return -1, fmt.Errorf("empty deque")
	}

	d.front++
	return d.data[d.front], nil
}

func (d *ArrayDeque) DequeueLast() (int, error) {
	// empty
	if d.front == d.rear {
		return -1, fmt.Errorf("empty deque")
	}
	v := d.data[d.rear]
	d.rear--
	return v, nil
}

func (d *ArrayDeque) PeekFirst() (int, error) {
	// empty
	if d.front == d.rear {
		return -1, fmt.Errorf("empty deque")
	}

	f := d.front
	f++
	return d.data[f], nil
}

func (d *ArrayDeque) PeekLast() (int, error) {
	// empty
	if d.front == d.rear {
		return -1, fmt.Errorf("empty deque")
	}

	return d.data[d.rear], nil
}

func (d *ArrayDeque) Empty() bool {
	return d.front == d.rear
}

func (d *ArrayDeque) Size() int {
	switch d.front {
	case d.rear:
		return 0
	case -1:
		return d.rear + 1
	default:
		return d.rear - d.front + 1
	}
}
