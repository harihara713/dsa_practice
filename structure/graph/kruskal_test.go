package graph

import "testing"

func TestKruskalMST(t *testing.T) {
	edges := []Edge{
		{Start: 1, End: 2, Weight: 28}, {Start: 1, End: 6, Weight: 10}, {Start: 2, End: 3, Weight: 16}, {Start: 2, End: 7, Weight: 14}, {Start: 3, End: 4, Weight: 12},
		{Start: 4, End: 5, Weight: 22}, {Start: 4, End: 7, Weight: 18}, {Start: 5, End: 6, Weight: 25}, {Start: 5, End: 7, Weight: 24},
	}

	wantedCost := 99
	cost := KruskalMST(edges, 7)

	if wantedCost != cost {
		t.Fatalf("Expected minimum cost = %d, Got minimum cost = %d", wantedCost, cost)
	}
}
