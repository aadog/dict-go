package dict

import "sync"

// DictList holds the set of elements in a slice
//
// It also stores the total number of elements currently present in the list, so length of the list is given in O(1)
//
// It uses sync.RWMutex to support concurrent access, all read operations acquires a RLock, and all write operatiors acquires a Lock()
type DictList struct {
	elements []*Dict
	size     int
	mtx      sync.RWMutex
	sync     bool
}

// New creates a new empty list and return it's reference.
func NewDictList() *DictList {
	return &DictList{sync: true}
}

// Len returns the number of elements in list
func (al *DictList) Len() int {
	al.rlock()
	defer al.runlock()

	return al.len()
}

// Front returns the first element of list or nil if the list is empty
func (al *DictList) Front() (*Dict, bool) {
	al.rlock()
	defer al.runlock()

	return al.front()
}

// Back returns the last element of the list or nil if the list is empty
func (al *DictList) Back() (interface{}, bool) {
	al.rlock()
	defer al.runlock()

	return al.back()
}

// PushFront inserts a new element with value v at the front of the list
func (al *DictList) PushFront(v *Dict) {
	al.lock()
	defer al.unlock()

	al.pushFront(v)
}

// PushBack inserts a new element with value v at the front of the list
func (al *DictList) PushBack(v *Dict) {
	al.lock()
	defer al.unlock()

	al.pushBack(v)
}

// Set inserts a new element with value v at the given index i.
// if index i is out of bound, it returns false, otherwise true
func (al *DictList) Set(i int, v *Dict) (ok bool) {
	al.lock()
	defer al.unlock()

	return al.set(i, v)
}

// Get ..
func (al *DictList) Get(i int) (v *Dict, ok bool) {
	al.rlock()
	defer al.runlock()

	return al.get(i)
}

// Insert value v at index i
func (al *DictList) Insert(i int, v *Dict) (ok bool) {
	al.lock()
	defer al.unlock()

	return al.insert(i, v)
}

// Remove the element at given index i. Returns true if element was removed otherwise false.
func (al *DictList) Remove(i int) (v interface{}, ok bool) {
	al.lock()
	defer al.unlock()

	return al.remove(i)
}

// Clear the list
func (al *DictList) Clear() {
	al.lock()
	defer al.unlock()

	al.clear()
}

// PushBackList inserts a copy of an other list at the back of list l.
// The lists l and other may be the same. They must not be nil.
func (al *DictList) PushBackList(l *DictList) {
	al.lock()
	defer al.unlock()

	al.pushBackList(l)
}

// PushFrontList inserts a copy of an other list at the front of list l.
// The lists l and other may be the same. They must not be nil.
func (al *DictList) PushFrontList(l *DictList) {
	al.lock()
	defer al.unlock()

	al.pushFrontList(l)
}

// Contains returns true if the given value exists in the list, otherwise false
func (al *DictList) Contains(v *Dict) bool {
	al.rlock()
	defer al.runlock()

	return al.contains(v)
}

// IndexOf returns the index of the given value v if it exists, otherwise it returns -1
func (al *DictList) IndexOf(v *Dict) int {
	al.rlock()
	defer al.runlock()

	return al.indexOf(v)
}

// Values returns all the values in the list as a slice
func (al *DictList) Values() []*Dict {
	al.rlock()
	defer al.runlock()

	return al.values()
}

// Clone creates a shallow copy and returns the reference
func (al *DictList) Clone() *DictList {
	al.rlock()
	defer al.runlock()

	return al.clone()
}

// Swap two values at two given indexes
func (al *DictList) Swap(a, b int) bool {
	al.rlock()
	defer al.runlock()

	return al.swap(a, b)
}

// String returns the string representation of the list
func (al *DictList) String() string {
	al.rlock()
	defer al.runlock()

	return al.string()
}

func (al *DictList) lock() {
	if al.sync {
		al.mtx.Lock()
	}
}

func (al *DictList) unlock() {
	if al.sync {
		al.mtx.Unlock()
	}
}

func (al *DictList) rlock() {
	if al.sync {
		al.mtx.RLock()
	}
}

func (al *DictList) runlock() {
	if al.sync {
		al.mtx.RUnlock()
	}
}
