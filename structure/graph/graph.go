package graph

import (
	"fmt"
	"slices"

	"github.com/harry713j/dsa_practice/constraints"
)

// Implementation of graph using adjacency list
type Graph[T constraints.Ordered] struct {
	adj      map[T][]T
	directed bool // indicate if the graph is a directed graph or not
	vertices int  // size of the vertices
}

func New[T constraints.Ordered](directed bool) *Graph[T] {
	return &Graph[T]{
		adj:      make(map[T][]T, 0),
		directed: directed,
	}
}

// Add a new vertex to the graph if it is not exists
func (g *Graph[T]) AddVertex(v T) {
	if _, ok := g.adj[v]; !ok {
		g.adj[v] = []T{}
		g.vertices++
	}
}

// Add a edge between two vertices
func (g *Graph[T]) AddEdge(from, to T) {
	g.AddVertex(from)
	g.AddVertex(to)

	if !slices.Contains(g.adj[from], to) {
		g.adj[from] = append(g.adj[from], to)
	}

	// if the graph is undirected
	if !g.directed {
		g.adj[to] = append(g.adj[to], from)
	}
}

// Breadth-First Search
func (g *Graph[T]) BFS(v T) []T {
	queue := make([]T, 0)
	visited := make(map[T]bool, g.vertices)
	res := make([]T, 0, g.vertices)

	if _, exist := g.adj[v]; !exist {
		return res
	}

	res = append(res, v)
	visited[v] = true
	queue = append(queue, v)

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for _, e := range g.adj[u] {
			if !visited[e] {
				res = append(res, e)
				visited[e] = true
				queue = append(queue, e)
			}
		}
	}

	return res
}

// print
func (g *Graph[T]) print() {
	for k, v := range g.adj {
		fmt.Printf("%v: [", k)
		for _, e := range v {
			fmt.Printf("%v ", e)
		}
		fmt.Print("]")
		fmt.Println()
	}
}

// Depth-First Search
func (g *Graph[T]) DFS(v T) []T {
	if _, exist := g.adj[v]; !exist {
		return []T{}
	}

	res := make([]T, 0)
	visited := make(map[T]bool, g.vertices)

	g.dfsHelper(visited, &res, v)
	return res
}

func (g *Graph[T]) dfsHelper(visited map[T]bool, res *[]T, v T) {
	if visited[v] {
		return
	}

	*res = append(*res, v)
	visited[v] = true

	for _, e := range g.adj[v] {
		if !visited[e] {
			g.dfsHelper(visited, res, e)
		}
	}
}
