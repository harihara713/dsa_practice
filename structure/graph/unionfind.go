package graph

// Disjoint Set
type UnionFind struct {
	parent []int
	// we can also maintain a rank array separately
}

func NewUnionFind(s int) *UnionFind {
	parent := make([]int, s)

	for i := 0; i < s; i++ {
		parent[i] = -1 // -1 means every element is its own parent
	}

	return &UnionFind{parent: parent}
}

// Return parent of the element
func (u *UnionFind) Find(q int) int {
	if u.parent[q] <= 0 {
		return q
	}

	u.parent[q] = u.Find(u.parent[q])
	return u.parent[q]
}

// Perform union
func (u *UnionFind) Union(p, q int) {
	rootP := u.Find(p)
	rootQ := u.Find(q)

	// if the root are same means both the elements belong to same set, if we union them then it will form a cycle
	if rootP == rootQ {
		return
	}

	if u.parent[rootP] < u.parent[rootQ] {
		u.parent[rootP] = u.parent[rootP] + u.parent[rootQ]
		u.parent[rootQ] = rootP
	} else {
		u.parent[rootQ] = u.parent[rootP] + u.parent[rootQ]
		u.parent[rootP] = rootQ
	}
}
