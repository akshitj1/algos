package ds

type Queue struct {
	els []int
}

func NewQueue() *Queue {
	return &Queue{make([]int, 0)}
}

func (q *Queue) Tail() (int, bool) {
	if q.Empty() {
		return 0, false
	}
	return q.els[0], true
}

func (q *Queue) Enqueue(x int) {
	q.els = append(q.els, x)
}

func (q *Queue) Dequeue() (int, bool) {
	if q.Empty() {
		return 0, false
	}
	el := q.els[0]
	q.els = q.els[1:]
	return el, true
}

func (q *Queue) Empty() bool {
	return len(q.els) == 0
}

func (q *Queue) Len() int {
	return len(q.els)
}
