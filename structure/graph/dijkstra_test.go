package graph

import (
	"math"
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	cost := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 5, 0, 0, 0},
		{0, 1, 0, 3, 10, 8, 0},
		{0, 5, 3, 0, 0, 2, 0},
		{0, 0, 10, 0, 0, 3, 2},
		{0, 0, 8, 2, 3, 0, 7},
		{0, 0, 0, 0, 2, 7, 0},
	}

	for _, group := range cost {
		for i := 1; i <= 6; i++ {
			if group[i] == 0 {
				group[i] = math.MaxInt
			}
		}
	}

	wantShortestPath := []int{0, 0, 1, 4, 9, 6, 11}
	shortestPath := Dijkstra(cost, 6, 1)

	if !reflect.DeepEqual(wantShortestPath, shortestPath) {
		t.Fatalf("Dijkstra's Shortest Path Algorithm:\nExpecting SP = %v, Got SP = %v", wantShortestPath, shortestPath)
	}
}
