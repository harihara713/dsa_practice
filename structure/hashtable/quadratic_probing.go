package hashtable

var defaultHtqpCapacity = 10 // default hash table capacity

// Hash table with quadratic probing
type HTQP struct {
	cap   int
	size  int
	table []int
}

// Hash table using linear probing with default capacity 10
func NewHTQP() *HTQP {
	return &HTQP{
		cap:   defaultHtlpCapacity,
		table: make([]int, defaultHtlpCapacity),
	}
}

// Insert the key in the Hash table
func (h *HTQP) Insert(key int) {
	idx := h.hash(key)

	// check if there is already a key present or not
	if h.table[idx] == 0 {
		h.table[idx] = key
	} else {
		idx = h.probe(key)
		h.table[idx] = key
	}

	h.size++
	loadingFactor := float32(h.size) / float32(h.cap)
	// if size is greater than the loading factor = 0.75
	if loadingFactor > 0.5 {
		// resize the table and re-hash
		h.resize()
	}
}

// Returns the Key if exists otherwise return -1
func (h *HTQP) Get(key int) int {
	idx, ok := h.getKeyIndex(key)
	if !ok {
		return -1
	}

	return h.table[idx]
}

// Return true if the key is exists in the Hash table, if not then returns false
func (h *HTQP) Contains(key int) bool {
	_, ok := h.getKeyIndex(key)
	return ok
}

// Delete and return the key if presents, if not then returns -1
func (h *HTQP) Delete(key int) int {
	idx, ok := h.getKeyIndex(key)
	if !ok {
		return -1
	}

	val := h.table[idx]
	// after delete re-insert the keys
	h.table[idx] = 0

	if idx != h.hash(key) {
		oldTable := h.table
		h.table = make([]int, h.cap)
		h.size = 0

		for _, k := range oldTable {
			if k != 0 {
				h.Insert(k)
			}
		}

	}

	return val
}

func (h *HTQP) hash(key int) int {
	return key % h.cap
}

func (h *HTQP) probe(key int) int {
	idx := 0
	for h.table[h.hash(key)+idx*idx] != 0 {
		idx++
	}

	return h.hash(key) + idx*idx
}

func (h *HTQP) resize() {
	oldTable := h.table
	h.cap <<= 1
	h.table = make([]int, h.cap)
	h.size = 0

	for _, k := range oldTable {
		if k != 0 {
			h.Insert(k)
		}
	}
}

func (h *HTQP) getKeyIndex(key int) (int, bool) {
	idx := 0
	for h.table[h.hash(key)+idx*idx] != 0 {
		if key == h.table[h.hash(key)+idx*idx] {
			return h.hash(key) + idx*idx, true
		}
		idx++
	}

	return -1, false
}
