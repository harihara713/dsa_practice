package heap

import (
	"github.com/harry713j/dsa_practice/constraints"
)

// Max Heap
type Heap[T constraints.Ordered] struct {
	data []T
	size int // current size of the Heap
}

// Create a Heap of size 16
func New[T constraints.Ordered]() *Heap[T] {
	return &Heap[T]{
		data: make([]T, 1, 16),
		size: 0,
	}
}

// Insert an element to the Heap
func (h *Heap[T]) Insert(element T) {
	h.data = append(h.data, element)
	h.size++
	h.shiftUp(h.size)
}

// Delete the root of the Heap which is happen to be the max or highest element in the Heap
func (h *Heap[T]) Delete() T {
	if h.Empty() {
		panic("Heap is empty")
	}

	return h.shiftDown(h.size)
}

// Top returns the root of element of the heap, (Max element of the heap)
func (h *Heap[T]) Top() T {
	if h.Empty() {
		panic("Heap is empty")
	}

	return h.data[1]
}

// Empty returns true if the heap is empty otherwise returns false
func (h *Heap[T]) Empty() bool {
	return h.size == 0
}

// Adjust the element in the Heap in its correct position
func (h *Heap[T]) shiftUp(n int) {
	temp := h.data[n]
	i := n

	for i > 1 && temp > h.data[i/2] {
		h.data[i] = h.data[i/2]
		i = i / 2
	}

	h.data[i] = temp
}

func (h *Heap[T]) shiftDown(n int) T {
	val := h.data[1]
	i := 1
	j := 2 * 1
	h.data[1] = h.data[n]

	for j < n {
		if h.data[j] < h.data[j+1] {
			j = j + 1
		}

		if h.data[i] < h.data[j] {
			// swap the value
			h.data[i], h.data[j] = h.data[j], h.data[i]
			i = j
			j *= 2
		} else {
			break
		}
	}

	h.data[n] = val
	h.size--
	return val
}
