package linkedlist

import "fmt"

type singleNode[T any] struct {
	value T
	next  *singleNode[T]
}

type SinglyLinkedList[T any] struct {
	head *singleNode[T] // pointing to the head of the list
	size int            // current size of the linked list
}

var ErrPositionOutOfBound = fmt.Errorf("position out of bound")

func newSingleNode[T any](value T) *singleNode[T] {
	return &singleNode[T]{
		value: value,
		next:  nil,
	}
}

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		head: nil,
		size: 0,
	}
}

// Append a new node to the end of the singly linked list
func (l *SinglyLinkedList[T]) Append(value T) {
	// head could be null
	node := newSingleNode[T](value)

	if l.head == nil {
		l.head = node
		l.size++
		return
	}
	// find the last node
	current := l.head

	for current.next != nil {
		current = current.next
	}

	current.next = node
	l.size++
}

// Insert the node with value to the specified position in the singly linked list, if the
// position is out of bound then it will return an error value otherwise nil
func (l *SinglyLinkedList[T]) Insert(pos int, value T) error {
	if pos < 0 || pos > l.size {
		return ErrPositionOutOfBound
	}
	// if head at nil
	if l.head == nil {
		l.Append(value)
		return nil
	}

	node := newSingleNode[T](value)
	current := l.head

	if pos == 0 {
		node.next = l.head
		l.head = node
		l.size++
		return nil
	}

	// move the head to the previous position where we want to insert the node
	for i := 0; i < pos-1; i++ {
		current = current.next
	}

	node.next = current.next
	current.next = node
	l.size++
	return nil
}

// Return the value of the node at the specified position in the list,
// if the position is out of bound it will return an error with value -1
func (l *SinglyLinkedList[T]) Get(pos int) (T, error) {
	if pos < 0 || pos > l.size-1 {
		var z T
		return z, ErrPositionOutOfBound
	}

	cur := l.head
	for i := 0; i < pos; i++ {
		cur = cur.next
	}

	return cur.value, nil
}

// Delete the node at the specified position, if the position is out of bound then it returns the error
func (l *SinglyLinkedList[T]) Delete(pos int) error {
	if pos < 0 || pos > l.size-1 {
		return ErrPositionOutOfBound
	}

	if l.head == nil {
		return nil
	}

	// if the position is head node
	if pos == 0 {
		temp := l.head
		l.head = temp.next

		// cut the link
		temp.next = nil
		l.size--
		return nil
	}

	cur := l.head
	for i := 0; i < pos-1; i++ {
		cur = cur.next
	}

	cur.next = cur.next.next
	l.size--
	return nil
}

// Return the size of the linked list
func (l *SinglyLinkedList[T]) Size() int {
	return l.size
}

// Print the content
func (l *SinglyLinkedList[T]) String() string {
	cur := l.head
	out := ""

	for cur != nil {
		out += fmt.Sprintf("%v -> ", cur.value)
		cur = cur.next
	}

	out += "nil"
	return out
}
