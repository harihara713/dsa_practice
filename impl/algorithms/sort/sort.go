package sort

import "github.com/harry713j/dsa_practice/constraints"

// Bubble Sort Algorithm, Time: O(n^2)
func BubbleSort[T constraints.Ordered](arr []T) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				// swap
				swap(arr, j, j+1)
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}

// Insertion Sort Algorithm, Time: O(n^2)
func InsertionSort[T constraints.Ordered](arr []T) {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				swap(arr, j, j-1)
			} else {
				break
			}
		}
	}
}

func swap[T constraints.Ordered](arr []T, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
