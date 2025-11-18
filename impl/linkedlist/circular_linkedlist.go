package linkedlist

import "fmt"

type circularNode struct {
	value int
	next  *circularNode
}

type CircularLinkedList struct {
	head *circularNode
	size int
}

func newCircularNode(val int) *circularNode {
	return &circularNode{
		value: val,
		next:  nil,
	}
}

func NewCircularLinkedList() *CircularLinkedList {
	return &CircularLinkedList{
		head: nil,
		size: 0,
	}
}

func (l *CircularLinkedList) Append(val int) {
	node := newCircularNode(val)

	if l.head == nil {
		l.head = node
		l.head.next = l.head
		l.size++
		return
	}

	curr := l.head
	curr = curr.next

	for curr.next != l.head {
		curr = curr.next
	}

	curr.next = node
	node.next = l.head
	l.size++
}

func (l *CircularLinkedList) Insert(pos, val int) error {
	if pos < 0 || pos > l.size {
		return ErrPositionOutOfBound
	}

	if l.head == nil {
		l.Append(val)
		return nil
	}

	node := newCircularNode(val)
	curr := l.head
	// pos == 0
	if pos == 0 {
		curr = curr.next
		for curr.next != l.head {
			curr = curr.next
		}

		curr.next = node
		node.next = l.head
		l.head = node
		l.size++
		return nil
	}

	for i := 0; i < pos-1; i++ {
		curr = curr.next
	}

	node.next = curr.next
	curr.next = node
	l.size++
	return nil
}

func (l *CircularLinkedList) Get(pos int) (int, error) {
	if pos < 0 || pos >= l.size {
		return -1, ErrPositionOutOfBound
	}

	if l.head == nil {
		return -1, fmt.Errorf("head is nil")
	}

	curr := l.head
	for i := 0; i < pos; i++ {
		curr = curr.next
	}

	return curr.value, nil
}

func (l *CircularLinkedList) Delete(pos int) (int, error) {
	if pos < 0 || pos >= l.size {
		return -1, ErrPositionOutOfBound
	}

	if l.head == nil {
		return -1, fmt.Errorf("head is nil")
	}

	curr := l.head
	if pos == 0 {
		curr = curr.next
		for curr.next != l.head {
			curr = curr.next
		}

		remove := l.head
		curr.next = l.head.next
		l.head = l.head.next
		l.size--
		return remove.value, nil
	}

	for i := 0; i < pos-1; i++ {
		curr = curr.next
	}

	remove := curr.next
	curr.next = curr.next.next
	l.size--
	return remove.value, nil
}

func (l *CircularLinkedList) Size() int {
	return l.size
}
