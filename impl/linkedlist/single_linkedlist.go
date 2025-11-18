package linkedlist

import "fmt"

type singleNode struct {
	value int
	next  *singleNode
}

type SinglyLinkedList struct {
	head *singleNode // pointing to the head of the list
	size int         // current size of the linked list
}

var ErrPositionOutOfBound = fmt.Errorf("position out of bound")

func newSingleNode(value int) *singleNode {
	return &singleNode{
		value: value,
		next:  nil,
	}
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		head: nil,
		size: 0,
	}
}

// Append a new node to the end of the singly linked list
func (l *SinglyLinkedList) Append(value int) {
	// head could be null
	node := newSingleNode(value)

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
func (l *SinglyLinkedList) Insert(pos int, value int) error {
	if pos < 0 || pos > l.size {
		return ErrPositionOutOfBound
	}
	// if head at nil
	if l.head == nil {
		l.Append(value)
		return nil
	}

	node := newSingleNode(value)
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
func (l *SinglyLinkedList) Get(pos int) (int, error) {
	if pos < 0 || pos > l.size-1 {
		return -1, ErrPositionOutOfBound
	}

	cur := l.head
	for i := 0; i < pos; i++ {
		cur = cur.next
	}

	return cur.value, nil
}

// Delete the node at the specified position, if the position is out of bound then it returns the error
func (l *SinglyLinkedList) Delete(pos int) error {
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
func (l *SinglyLinkedList) Size() int {
	return l.size
}

// Print the content
func (l *SinglyLinkedList) String() string {
	cur := l.head
	out := ""

	for cur != nil {
		out += fmt.Sprintf("%d -> ", cur.value)
		cur = cur.next
	}

	out += "nil"
	return out
}
