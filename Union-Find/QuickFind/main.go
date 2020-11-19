package main

// QuickFind ...
type QuickFind struct {
	id []int
}

// New ...
func New(n int) *QuickFind {
	return &QuickFind{
		id: make([]int, n),
	}
}

// Connected ...
func (qf *QuickFind) Connected(p, q int) bool {
	return qf.id[p] == qf.id[q]
}

// Union ...
func (qf *QuickFind) Union(p, q int) {
	pid := qf.id[p] // parent of p
	qid := qf.id[q] // parent of q

	for index, value := range qf.id {
		if value == pid { // Set parent as q for all the id where parent is of p
			qf.id[index] = qid
		}
	}

}

// Init - Intialise the parent array as same element
func (qf *QuickFind) Init() {
	for index := range qf.id {
		qf.id[index] = index
	}
}

func main() {
	uf := New(10)
	uf.Init()
}
