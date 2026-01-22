package graph

import (
	"reflect"
	"slices"
	"sort"
	"testing"

	"github.com/harry713j/dsa_practice/constraints"
)

func sorted[T constraints.Ordered](slice []T) []T {
	s := make([]T, len(slice))
	copy(s, slice)
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return s
}

// TestUndirectedGraph verifies basic behavior of an undirected graph
func TestUndirectedGraph(t *testing.T) {
	g := New[int](false) // undirected

	// Add edges (automatically adds vertices)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)

	// Test: vertex count
	if g.vertices != 5 {
		t.Errorf("expected 5 vertices, got %d", g.vertices)
	}

	tests := []struct {
		v        int
		expected []int
	}{
		{1, []int{2, 3}},
		{2, []int{1, 4}},
		{3, []int{1, 4}},
		{4, []int{2, 3, 5}},
		{5, []int{4}},
	}

	for _, tt := range tests {
		got := sorted(g.adj[tt.v])
		want := sorted(tt.expected)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("neighbors of %d: got %v, want %v", tt.v, got, want)
		}
	}

	g.AddEdge(1, 2)
	if len(g.adj[1]) != 2 { // still only 2 neighbors
		t.Errorf("duplicate edge added - neighbors of 1: %v", g.adj[1])
	}

	// Test: adding isolated vertex
	g.AddVertex(10)
	if g.vertices != 6 {
		t.Errorf("after AddVertex(10): expected 6 vertices, got %d", g.vertices)
	}
	if len(g.adj[10]) != 0 {
		t.Errorf("new vertex 10 should have no neighbors, got %v", g.adj[10])
	}
}

func TestDirectedGraph(t *testing.T) {
	g := New[string](true) // directed

	// Build a small directed graph
	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")
	g.AddEdge("C", "D")
	g.AddEdge("D", "E")
	g.AddEdge("E", "A") // cycle

	// Test: vertex count
	if g.vertices != 5 {
		t.Errorf("expected 5 vertices, got %d", g.vertices)
	}

	// Test: directed edges (no reverse unless explicitly added)
	tests := []struct {
		v        string
		expected []string
	}{
		{"A", []string{"B", "C"}},
		{"B", []string{"D"}},
		{"C", []string{"D"}},
		{"D", []string{"E"}},
		{"E", []string{"A"}},
	}

	for _, tt := range tests {
		got := sorted(g.adj[tt.v])
		want := sorted(tt.expected)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("outgoing neighbors of %q: got %v, want %v", tt.v, got, want)
		}
	}

	// Test: reverse edges should NOT exist automatically
	if slices.Contains(g.adj["B"], "A") {
		t.Error("directed graph should not have reverse edge B→A")
	}

	if slices.Contains(g.adj["D"], "B") {
		t.Error("directed graph should not have reverse edge D→B")
	}

	// Test: adding the same directed edge multiple times
	g.AddEdge("A", "B")
	if len(g.adj["A"]) == 3 {
		t.Errorf("expected no duplicate edge to be added, got neighbors: %v", g.adj["A"])
	}
}

func TestBFSGraph(t *testing.T) {
	g := New[int](false) // undirected

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)

	tests := []struct {
		name   string
		vertex int
		want   []int
	}{
		{
			name:   "Vertex not present in graph",
			vertex: 7,
			want:   []int{},
		},
		{
			name:   "Vertex start with 1",
			vertex: 1,
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "Vertex start with 5",
			vertex: 5,
			want:   []int{5, 4, 2, 3, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := g.BFS(tt.vertex)

			if !reflect.DeepEqual(res, tt.want) {
				t.Fatalf("BFS: expected = %v, got = %v", tt.want, res)
			}
		})
	}
}

func TestDFSGraph(t *testing.T) {
	g := New[int](false) // undirected

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)

	tests := []struct {
		name   string
		vertex int
		want   []int
	}{
		{
			name:   "Vertex not present in graph",
			vertex: 7,
			want:   []int{},
		},
		{
			name:   "Vertex start with 1",
			vertex: 1,
			want:   []int{1, 2, 4, 3, 5},
		},
		{
			name:   "Vertex start with 5",
			vertex: 5,
			want:   []int{5, 4, 2, 1, 3},
		},
	}

	g.print()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := g.DFS(tt.vertex)

			if !reflect.DeepEqual(res, tt.want) {
				t.Fatalf("DFS: expected = %v, got = %v", tt.want, res)
			}
		})
	}
}
