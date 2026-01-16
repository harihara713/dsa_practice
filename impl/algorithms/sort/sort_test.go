package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "Unsorted",
			arr:  []int{40, 30, 5, 15, 11, 13},
			want: []int{5, 11, 13, 15, 30, 40},
		},
		{
			name: "Already sorted",
			arr:  []int{1, 3, 6, 15, 23},
			want: []int{1, 3, 6, 15, 23},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.arr)

			if !reflect.DeepEqual(tt.arr, tt.want) {
				t.Fatalf("expected = %v, got = %v\n", tt.want, tt.arr)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "Unsorted",
			arr:  []int{40, 30, 5, 15, 11, 13},
			want: []int{5, 11, 13, 15, 30, 40},
		},
		{
			name: "Already sorted",
			arr:  []int{1, 3, 6, 15, 23},
			want: []int{1, 3, 6, 15, 23},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertionSort(tt.arr)

			if !reflect.DeepEqual(tt.arr, tt.want) {
				t.Fatalf("expected = %v, got = %v\n", tt.want, tt.arr)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "Unsorted",
			arr:  []int{40, 30, 5, 15, 11, 13},
			want: []int{5, 11, 13, 15, 30, 40},
		},
		{
			name: "Already sorted",
			arr:  []int{1, 3, 6, 15, 23},
			want: []int{1, 3, 6, 15, 23},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectionSort(tt.arr)

			if !reflect.DeepEqual(tt.arr, tt.want) {
				t.Fatalf("expected = %v, got = %v\n", tt.want, tt.arr)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "Unsorted",
			arr:  []int{40, 30, 5, 15, 11, 13},
			want: []int{5, 11, 13, 15, 30, 40},
		},
		{
			name: "Already sorted",
			arr:  []int{1, 3, 6, 15, 23},
			want: []int{1, 3, 6, 15, 23},
		},
		{
			name: "Descending Sorted",
			arr:  []int{50, 40, 30, 20, 10},
			want: []int{10, 20, 30, 40, 50},
		},
		{
			name: "Unsorted mixed",
			arr:  []int{50, 20, 30, 10, 80, 90, 70, 60},
			want: []int{10, 20, 30, 50, 60, 70, 80, 90},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.arr)

			if !reflect.DeepEqual(tt.arr, tt.want) {
				t.Fatalf("expected = %v, got = %v\n", tt.want, tt.arr)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "Unsorted",
			arr:  []int{40, 30, 5, 15, 11, 13},
			want: []int{5, 11, 13, 15, 30, 40},
		},
		{
			name: "Already sorted",
			arr:  []int{1, 3, 6, 15, 23},
			want: []int{1, 3, 6, 15, 23},
		},
		{
			name: "Descending Sorted",
			arr:  []int{50, 40, 30, 20, 10},
			want: []int{10, 20, 30, 40, 50},
		},
		{
			name: "Unsorted mixed",
			arr:  []int{50, 20, 30, 10, 80, 90, 70, 60},
			want: []int{10, 20, 30, 50, 60, 70, 80, 90},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.arr)

			if !reflect.DeepEqual(tt.arr, tt.want) {
				t.Fatalf("expected = %v, got = %v\n", tt.want, tt.arr)
			}
		})
	}
}

func TestMergeSortIterative(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "Unsorted",
			arr:  []int{40, 30, 5, 15, 11, 13},
			want: []int{5, 11, 13, 15, 30, 40},
		},
		{
			name: "Already sorted",
			arr:  []int{1, 3, 6, 15, 23},
			want: []int{1, 3, 6, 15, 23},
		},
		{
			name: "Descending Sorted",
			arr:  []int{50, 40, 30, 20, 10},
			want: []int{10, 20, 30, 40, 50},
		},
		{
			name: "Unsorted mixed",
			arr:  []int{50, 20, 30, 10, 80, 90, 70, 60},
			want: []int{10, 20, 30, 50, 60, 70, 80, 90},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSortIterative(tt.arr)

			if !reflect.DeepEqual(tt.arr, tt.want) {
				t.Fatalf("expected = %v, got = %v\n", tt.want, tt.arr)
			}
		})
	}
}

func TestCountSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []uint
		want []uint
	}{
		{
			name: "Unsorted",
			arr:  []uint{40, 30, 5, 15, 11, 13},
			want: []uint{5, 11, 13, 15, 30, 40},
		},
		{
			name: "Already sorted",
			arr:  []uint{1, 3, 6, 15, 23},
			want: []uint{1, 3, 6, 15, 23},
		},
		{
			name: "Descending Sorted",
			arr:  []uint{50, 40, 30, 20, 10},
			want: []uint{10, 20, 30, 40, 50},
		},
		{
			name: "Unsorted mixed",
			arr:  []uint{50, 20, 30, 10, 80, 90, 70, 60},
			want: []uint{10, 20, 30, 50, 60, 70, 80, 90},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CountSort(tt.arr)

			if !reflect.DeepEqual(tt.arr, tt.want) {
				t.Fatalf("expected = %v, got = %v\n", tt.want, tt.arr)
			}
		})
	}
}

func TestBucketSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []uint
		want []uint
	}{
		{
			name: "Unsorted",
			arr:  []uint{40, 30, 5, 15, 11, 13},
			want: []uint{5, 11, 13, 15, 30, 40},
		},
		{
			name: "Already sorted",
			arr:  []uint{1, 3, 6, 15, 23},
			want: []uint{1, 3, 6, 15, 23},
		},
		{
			name: "Descending Sorted",
			arr:  []uint{50, 40, 30, 20, 10},
			want: []uint{10, 20, 30, 40, 50},
		},
		{
			name: "Unsorted mixed",
			arr:  []uint{50, 20, 30, 10, 80, 90, 70, 60},
			want: []uint{10, 20, 30, 50, 60, 70, 80, 90},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BucketSort(tt.arr)

			if !reflect.DeepEqual(tt.arr, tt.want) {
				t.Fatalf("expected = %v, got = %v\n", tt.want, tt.arr)
			}
		})
	}
}

func TestRadixSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []uint
		want []uint
	}{
		{
			name: "Unsorted",
			arr:  []uint{40, 30, 5, 15, 11, 13},
			want: []uint{5, 11, 13, 15, 30, 40},
		},
		{
			name: "Already sorted",
			arr:  []uint{1, 3, 6, 15, 23},
			want: []uint{1, 3, 6, 15, 23},
		},
		{
			name: "Descending Sorted",
			arr:  []uint{50, 40, 30, 20, 10},
			want: []uint{10, 20, 30, 40, 50},
		},
		{
			name: "Unsorted mixed",
			arr:  []uint{148, 199, 21, 43, 69, 13, 7, 85, 302},
			want: []uint{7, 13, 21, 43, 69, 85, 148, 199, 302},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RadixSort(tt.arr)

			fmt.Println(tt.arr)

			if !reflect.DeepEqual(tt.arr, tt.want) {
				t.Fatalf("expected = %v, got = %v\n", tt.want, tt.arr)
			}
		})
	}
}
