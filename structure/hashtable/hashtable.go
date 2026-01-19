package hashtable

import (
	"fmt"
	"hash/fnv"
)

var defaultCapacity int = 13

type Entry struct {
	key   any
	value any
}

type Hashtable struct {
	capacity int
	size     int
	table    []*Entry
}

func New() *Hashtable {
	return &Hashtable{
		capacity: defaultCapacity,
		table:    make([]*Entry, defaultCapacity),
	}
}

// Insert the key and value to the Hash table
func (ht *Hashtable) Put(key, value any) {
	idx, ok := ht.searchKey(key)
	if ok {
		// overwrite
		ht.table[idx].value = value
		return
	}

	idx = ht.probeIndex(key)
	ht.table[idx] = &Entry{key: key, value: value}
	ht.size++

	loadingFactor := float32(ht.size) / float32(ht.capacity)
	if loadingFactor > 0.5 {
		ht.resize()
	}
}

// Returns the value associate with the key if the key is not exists in the hash table it will return nil
func (ht *Hashtable) Get(key any) any {
	idx, ok := ht.searchKey(key)
	if !ok {
		return nil
	}

	return ht.table[idx].value
}

// Returns true if the key exists in the hash table and false if the key doesn't exist in the hash table
func (ht *Hashtable) Contains(key any) bool {
	_, ok := ht.searchKey(key)
	return ok
}

// Remove the entry containing the key if it exists in hash table and return the value otherwise returns nil
func (ht *Hashtable) Delete(key any) any {
	idx, ok := ht.searchKey(key)
	if !ok {
		return nil
	}

	val := ht.table[idx].value
	ht.table[idx] = nil

	if idx != ht.probe(key, 0) {
		oldTable := ht.table
		ht.table = make([]*Entry, ht.capacity)
		ht.size = 0

		for _, e := range oldTable {
			if e != nil {
				ht.Put(e.key, e.value)
			}
		}
	}

	return val
}

func (ht *Hashtable) hash1(key any) int {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%v", key)))

	hashVal := h.Sum32()
	return int(hashVal)
}

func (ht *Hashtable) hash2(key any) int {
	h := fnv.New32()
	h.Write([]byte(fmt.Sprintf("%v", key)))

	pNum := lastPrime(ht.capacity)
	return pNum - (int(h.Sum32()) % pNum)
}

func (ht *Hashtable) probeIndex(key any) int {
	i := 0
	for ht.table[ht.probe(key, i)] != nil {
		i++
	}

	return ht.probe(key, i)
}

func (ht *Hashtable) probe(key any, i int) int {
	return (ht.hash1(key) + i*ht.hash2(key)) % ht.capacity
}

func (ht *Hashtable) resize() {
	oldTable := ht.table
	ht.capacity <<= 1
	ht.table = make([]*Entry, ht.capacity)
	ht.size = 0

	for _, e := range oldTable {
		if e != nil {
			ht.Put(e.key, e.value)
		}
	}
}

func (ht *Hashtable) searchKey(key any) (int, bool) {
	i := 0
	for ht.table[ht.probe(key, i)] != nil {
		if key == ht.table[ht.probe(key, i)].key {
			return ht.probe(key, i), true
		}
		i++
	}

	return -1, false
}

// return max prime number in a range
func lastPrime(n int) int {
	for i := n; i >= 2; i-- {
		if prime(i) {
			return i
		}
	}

	return -1
}

func prime(n int) bool {
	if n < 2 {
		return false
	}

	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
