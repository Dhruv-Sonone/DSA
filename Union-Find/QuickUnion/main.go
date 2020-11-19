package main

// QuickUnion - Union operation is quick and Find operation is expensive
// Array stores the direct parent of the object

// QuickUnion ...
type QuickUnion struct {
	id []int
}

// New ...
func New(n int) *QuickUnion {
	return &QuickUnion{
		id: make([]int, n),
	}
}

// Init ...
func (qu *QuickUnion) Init() {
	for index := range qu.id {
		qu.id[index] = index
	}
}

// Connected ....
func (qu *QuickUnion) Connected(p, q int) bool {
	return qu.root(p) == qu.root(q)
}

// Union ...
func (qu *QuickUnion) Union(p, q int) {
	pid := qu.root(p)
	qid := qu.root(q)

	qu.id[pid] = qid

}

func (qu *QuickUnion) root(p int) int {
	index := p
	for qu.id[index] != index {
		index = qu.id[p]
	}

	return index
}

func main() {
	qu := New(10)
	qu.Init()
}
