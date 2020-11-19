package main

// QuickUnion ...
type QuickUnion struct {
	id   []int
	size []int
}

// New ...
func New(n int) *QuickUnion {
	return &QuickUnion{
		id:   make([]int, n),
		size: make([]int, n),
	}
}

// Init ...
func (qu *QuickUnion) Init() {
	for index := range qu.id {
		qu.id[index] = index
		qu.size[index] = 1
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

	if qu.size[pid] < qu.size[qid] {
		qu.id[pid] = qid
		qu.size[pid] = qu.size[pid] + qu.size[qid]
	} else {
		qu.id[qid] = pid
		qu.size[qid] = qu.size[pid] + qu.size[qid]
	}

}

func (qu *QuickUnion) root(p int) int {
	index := p
	for qu.id[index] != index {
		index = qu.id[p]
	}

	return index
}

func main() {

}
