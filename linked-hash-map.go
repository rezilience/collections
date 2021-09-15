package collections

/* ****** WIP ****** */

type (
	linkedHashSet struct {
		m    map[int]*node
		head *node
		tail *node
		size int
	}
)

func NewLinkedHashSet() *linkedHashSet {
	l := &linkedHashSet{}
	l.m = make(map[int]*node)
	return l
}

func (l *linkedHashSet) Add(e int) bool {
	if _, exists := l.m[e]; exists {
		return false
	}
	n := &node{e: e}
	l.m[e] = n
	l.size++
	if l.tail == nil { // Map is empty.
		l.head, l.tail = n, n
	} else { // Map is non-empty. Add at the tail.
		l.tail.next = n
		l.tail = n
	}
	return true
}

func (l *linkedHashSet) Contains(e int) bool {
	_, exists := l.m[e]
	return exists
}

func (l *linkedHashSet) First() *int {
	if l.head == nil {
		return nil
	} else {
		i := l.head.e.(int)
		return &i
	}
}

func (l *linkedHashSet) Last() *int {
	if l.tail == nil {
		return nil
	} else {
		i := l.tail.e.(int)
		return &i
	}
}

func (l *linkedHashSet) Size() int {
	return l.size
}
