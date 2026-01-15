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

// Selection Sort Algorithm, Time: O(n^2)
func SelectionSort[T constraints.Ordered](arr []T) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		k := i
		for j := i; j < n; j++ {
			if arr[j] < arr[k] {
				k = j
			}
		}

		// swap
		swap(arr, i, k)
	}
}

// Quick Sort Algorithm, Worst Time: O(n^2) when array is sorted, Avg Time: O(nlogn)
func QuickSort[T constraints.Ordered](arr []T) {
	n := len(arr)
	quickSort(arr, 0, n-1)
}

func quickSort[T constraints.Ordered](arr []T, low, high int) {
	if high-low+1 <= 1 {
		return
	}

	p := partition(arr, low, high)
	quickSort(arr, low, p-1)
	quickSort(arr, p+1, high)
}

func partition[T constraints.Ordered](arr []T, l, h int) int {
	pivot := arr[l]
	i, j := l, h

	for i < j {
		for i <= h && arr[i] <= pivot {
			i++
		}

		for j >= l && arr[j] > pivot {
			j--
		}

		if i < j {
			swap(arr, i, j)
		}
	}

	// put pivot in its right position
	swap(arr, l, j)
	return j
}

func swap[T constraints.Ordered](arr []T, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
