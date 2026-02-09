package graph

import (
	"math"
	"testing"
)

func TestPrimMST(t *testing.T) {
	// fill up the graph weight or cost value
	g := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 28, 0, 0, 0, 10, 0},
		{0, 28, 0, 16, 0, 0, 0, 14},
		{0, 0, 16, 0, 12, 0, 0, 0},
		{0, 0, 0, 12, 0, 22, 0, 18},
		{0, 0, 0, 0, 22, 0, 25, 24},
		{0, 10, 0, 0, 0, 25, 0, 0},
		{0, 0, 14, 0, 18, 24, 0, 0},
	}

	// Initialising all to max int value
	for i := 1; i < len(g); i++ {
		for j := 1; j < len(g); j++ {
			if g[i][j] == 0 {
				g[i][j] = math.MaxInt
			}
		}
	}

	wantCost := 99
	cost := PrimMST(g, 7)

	if wantCost != cost {
		t.Fatalf("Prim's Algorithm:\nExpected Cost = %d, Got Cost = %d", wantCost, cost)
	}
}
