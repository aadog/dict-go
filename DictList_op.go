package dict

import "fmt"

func (al *DictList) len() int {
	return al.size
}

func (al *DictList) front() (*Dict, bool) {
	return al.get(0)
}

func (al *DictList) back() (*Dict, bool) {
	return al.get(al.len() - 1)
}

func (al *DictList) pushFront(v *Dict) {
	nal := make([]*Dict, al.size+1)
	copy(nal, []*Dict{v})
	copy(nal[1:], al.elements)
	al.elements = nal
	al.size++
}

func (al *DictList) pushBack(v *Dict) {
	al.elements = append(al.elements, v)
	al.size++
}

func (al *DictList) set(i int, v *Dict) bool {
	if al.withInRange(i) {
		al.elements[i] = v
		return true
	}
	return false
}

func (al *DictList) get(i int) (*Dict, bool) {
	if al.withInRange(i) {
		return al.elements[i], true
	}
	return nil, false
}

func (al *DictList) insert(i int, v *Dict) bool {
	if i >= 0 && i <= al.len() {
		al.elements = append(al.elements, v)
		copy(al.elements[i+1:], al.elements[i:])
		al.elements[i] = v
		al.size++

		return true
	}

	return false
}

func (al *DictList) remove(index int) (*Dict, bool) {
	if al.withInRange(index) {
		value := al.elements[index]

		newElems := make([]*Dict, al.size-1)

		copy(newElems, al.elements[:index])
		copy(newElems[index:], al.elements[index+1:])

		al.elements = newElems
		al.size--

		return value, true
	}

	return nil, false
}

func (al *DictList) clear() {
	al.elements = make([]*Dict, 0)
	al.size = 0
}

func (al *DictList) pushBackList(l *DictList) {
	al.elements = append(al.elements, l.Values()...)
	al.size += l.Len()
}

func (al *DictList) pushFrontList(l *DictList) {
	al.elements = append(l.Values(), al.elements...)
	al.size += l.Len()
}

func (al *DictList) contains(value *Dict) bool {
	if al.indexOf(value) == -1 {
		return false
	}

	return true
}

func (al *DictList) indexOf(value *Dict) int {
	for i, element := range al.elements {
		if element == value {
			return i
		}
	}

	return -1
}

func (al *DictList) values() []*Dict {
	return al.elements
}

func (al *DictList) clone() *DictList {
	nal := NewDictList()
	nal.pushFrontList(al)

	return nal
}

func (al *DictList) swap(a, b int) bool {
	if al.withInRange(a) && al.withInRange(b) {
		al.elements[a], al.elements[b] = al.elements[b], al.elements[a]
		return true
	}
	return false
}

func (al *DictList) withInRange(index int) bool {
	return index > -1 && index < al.size
}

func (al *DictList) string() string {
	return fmt.Sprintf("%v", al.elements)
}
