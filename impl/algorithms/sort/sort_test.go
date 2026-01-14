package sort

import (
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
				t.Fatalf("expected = %v, got = %v\n", tt.arr, tt.want)
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
				t.Fatalf("expected = %v, got = %v\n", tt.arr, tt.want)
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
				t.Fatalf("expected = %v, got = %v\n", tt.arr, tt.want)
			}
		})
	}
}
