package hashtable

// Hash Table with Chaining
type HTC struct {
	cap   int
	table [][]int
}

// Put, Search, Delete, hash function

func NewHTC() *HTC {
	h := &HTC{
		cap:   10,
		table: make([][]int, 10),
	}

	for i := range h.table {
		h.table[i] = make([]int, 0)
	}

	return h
}

// Insert the key to the Hash table
func (h *HTC) Insert(key int) {
	idx := h.hash(key)

	// insert the key in the index
	h.table[idx] = append(h.table[idx], key)
}

// Returns the index of the Key if present in Hash table, if not then returns -1
func (h *HTC) Get(key int) int {
	idx, ok := h.getKeyIndex(key)

	if !ok {
		return -1
	}

	return idx
}

// Returns true if the key exists in the Hash table, otherwise return false
func (h *HTC) Contains(key int) bool {
	_, ok := h.getKeyIndex(key)
	return ok
}

// Delete a key from the Hash table and return that value if that value exists in the Hash table,
// if not exists then it returns -1
func (h *HTC) Delete(key int) int {
	idx, ok := h.getKeyIndex(key)
	if !ok {
		return -1
	}

	// index of the key inside the bucket
	delIdx := -1
	for i, k := range h.table[idx] {
		if key == k {
			delIdx = i
			break
		}
	}

	val := h.table[idx][delIdx]

	h.table[idx] = append(h.table[idx][:delIdx], h.table[idx][delIdx+1:]...)
	return val
}

func (h *HTC) hash(key int) int {
	return key % h.cap
}

func (h *HTC) getKeyIndex(key int) (int, bool) {
	idx := h.hash(key)

	// search in the bucket
	for _, k := range h.table[idx] {
		if k == key {
			return idx, true
		}
	}

	return -1, false
}
