package main

import "fmt"

type DynamicArray struct {
	data []int // store the data
	size int   // current size of the array
}

func NewDynamicArray() *DynamicArray {
	return &DynamicArray{
		size: 0,
		data: make([]int, 5),
	}
}

// Append the value at the end of the dynamic array
func (d *DynamicArray) Append(value int) {
	if d.size == len(d.data) {
		d.resize()
	}

	d.data[d.size] = value
	d.size++
}

// Insert the value to a specific position, if the position is out of bound it return error
func (d *DynamicArray) Insert(pos int, value int) error {
	if pos > d.size || pos < 0 {
		return fmt.Errorf("position is out of bound")
	}

	if d.size == len(d.data) {
		d.resize()
	}

	if pos == d.size {
		d.data[d.size] = value
		d.size++
		return nil
	}

	// if inserting in-between, we have to shift to the right
	for i := d.size - 1; i >= pos; i-- {
		d.data[i+1] = d.data[i]
	}

	d.data[pos] = value
	d.size++
	return nil
}

// Set the previous value to current value in a specified position
func (d *DynamicArray) Set(pos int, value int) error {
	if pos > d.size || pos < 0 {
		return fmt.Errorf("position is out of bound")
	}

	if d.size == len(d.data) {
		d.resize()
	}

	if pos == d.size {
		d.data[pos] = value
		d.size++
	}

	d.data[pos] = value
	return nil
}

// Return the element at a position `pos` and return true if that position is valid else return false
func (d *DynamicArray) Get(pos int) (int, bool) {
	if pos >= d.size || pos < 0 {
		return -1, false
	}

	return d.data[pos], true
}

// Return current size of the DynamicArray
func (d *DynamicArray) Size() int {
	return d.size
}

// Delete element from a specified position pos, if the position is invalid it returns error
func (d *DynamicArray) Delete(pos int) error {
	if pos >= d.size || pos < 0 {
		return fmt.Errorf("position is out of bound")
	}

	if pos == d.size-1 {
		d.data[pos] = 0
		d.size--
		return nil
	}

	// deleting between the elements, so after deletion perform the left shift
	i := pos
	for i < d.size-1 {
		d.data[i] = d.data[i+1]
		i++
	}

	d.data[d.size-1] = 0
	d.size--
	return nil
}

func (d *DynamicArray) resize() {
	temp := make([]int, 2*len(d.data))
	copy(temp, d.data)

	d.data = temp
}
