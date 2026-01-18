package hashtable

import (
	"fmt"
	"hash/fnv"
)

// TODO: Do this with `Double Hashing` Technique

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

	idx = ht.hash(key)
	fmt.Printf("key = %v, value = %v, idx = %d\n", key, value, idx)
	if ht.table[idx] == nil {
		ht.table[idx] = &Entry{key: key, value: value}
	} else {
		idx = ht.probe(key)
		ht.table[idx] = &Entry{key: key, value: value}
	}

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

	fmt.Printf("idx = %d\n", idx)
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

	if idx != ht.hash(key) {
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

func (ht *Hashtable) hash(key any) int {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%v", key)))

	hashVal := h.Sum32()
	return int(hashVal) % ht.capacity
}

func (ht *Hashtable) probe(key any) int {
	idx := ht.hash(key)
	i := 0
	for ht.table[idx+i*i] != nil {
		i++
	}

	return idx + i*i
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
	idx := ht.hash(key)
	fmt.Printf("searchKey: idx = %d\n", idx)
	for ht.table[idx+i*i] != nil {
		fmt.Printf("searchKey: idx+i*i = %d\n", idx+i*i)
		if key == ht.table[idx+i*i].key {
			return idx + i*i, true
		}
		i++
	}

	return -1, false
}
