package sort

import (
	"fmt"

	"github.com/harry713j/dsa_practice/constraints"
	"github.com/harry713j/dsa_practice/impl/linkedlist"
)

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
			i++
			j--
		}
	}

	// put pivot in its right position
	swap(arr, l, j)
	return j
}

// Merge Sort Algorithm, Time: O(nlogn)
func MergeSort[T constraints.Ordered](arr []T) {
	n := len(arr)
	mergeSort(arr, 0, n-1)
}

func mergeSort[T constraints.Ordered](arr []T, low, high int) {
	if low >= high {
		return
	}

	mid := (low + high) / 2
	mergeSort(arr, low, mid)
	mergeSort(arr, mid+1, high)
	merge(arr, low, mid, high)
}

func merge[T constraints.Ordered](arr []T, low, mid, high int) {
	temp := make([]T, high-low+1)
	i, j, k := low, mid+1, 0

	for i <= mid && j <= high {
		if arr[i] < arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}

	for ; i <= mid; i++ {
		temp[k] = arr[i]
		k++
	}

	for ; j <= high; j++ {
		temp[k] = arr[j]
		k++
	}

	// copy the temp to arr
	for l := range temp {
		arr[low+l] = temp[l]
	}
}

func MergeSortIterative[T constraints.Ordered](arr []T) {
	n := len(arr)
	var p int

	for p = 2; p <= n; p *= 2 {
		for i := 0; i+p-1 < n; i = i + p {
			low := i
			high := i + p - 1
			mid := (low + high) / 2
			merge(arr, low, mid, high)
		}
	}

	if p/2 < n {
		merge(arr, 0, p/2-1, n-1)
	}
}

// Count Sort Algorithm, Time: O(n), Space: O(n)
func CountSort[T constraints.Unsigned](arr []T) {
	n := len(arr)
	max := maxNum(arr)
	count := make([]int, max+1)

	for i := 0; i < n; i++ {
		count[arr[i]]++
	}

	var j, k int
	for j < len(count) {
		if count[j] != 0 {
			arr[k] = T(j)
			count[j]--
			k++
		} else {
			j++
		}
	}
}

// Bucket Sort Algorithm, Time: O(n), Space: O(n)
func BucketSort[T constraints.Unsigned](arr []T) {
	n := len(arr)
	max := maxNum(arr)
	bucket := make([]linkedlist.SinglyLinkedList[T], max+1)

	// initialize
	for _, ll := range bucket {
		ll.Append(0)
	}

	for i := 0; i < n; i++ {
		bucket[arr[i]].Append(arr[i])
	}

	var j, k int
	for j < len(bucket) {
		var b linkedlist.SinglyLinkedList[T]
		for bucket[j] != b {
			val, err := bucket[j].Delete(0)
			if err != nil {
				fmt.Println("Error ", err)
				return
			}

			arr[k] = val
			k++
		}
		j++
	}
}

// Radix Sort Algorithm, Time: O(n), Space: O(n)
func RadixSort[T constraints.Unsigned](arr []T) {
	n := len(arr)
	max := maxNum(arr)
	bin := make([]linkedlist.SinglyLinkedList[T], 10)

	i := 1

	for T(i) < max {
		for j := 0; j < n; j++ {
			bin[(arr[j]/T(i))%10].Append(arr[j])
		}

		// put it to the arr
		var l int
		for k := range bin {
			var z linkedlist.SinglyLinkedList[T]
			for bin[k] != z {
				val, err := bin[k].Delete(0)
				if err != nil {
					fmt.Println("Error ", err)
					return
				}

				arr[l] = val
				l++
			}
		}

		i = i * 10
	}
}

// Shell Sort Algorithm, Time: O(nlogn),
// Based on the idea of Insertion Sort
func ShellSort[T constraints.Ordered](arr []T) {
	n := len(arr)
	gap := n / 2

	for gap > 0 {
		for i := 0; i+gap < n; i++ {
			if arr[i+gap] < arr[i] {
				swap(arr, i, i+gap)
				// also check previous elements
				j := i
				for j-gap >= 0 && arr[j] < arr[j-gap] {
					swap(arr, j-gap, j)
					j = j - gap
				}
			}
		}

		gap = gap / 2
	}
}

func swap[T constraints.Ordered](arr []T, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func maxNum[T constraints.Number](arr []T) T {
	max := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}
