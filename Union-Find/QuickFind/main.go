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
	pid := qf.id[p]
	qid := qf.id[q]

	for index, value := range qf.id {
		if value == pid {
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
