package ds

import "testing"

func TestQueue(t *testing.T) {
	q := &Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	t.Log(q)
	el1, _ := q.Dequeue()
	el2, _ := q.Dequeue()
	el3, _ := q.Dequeue()

	if !(el1 == 1 && el2 == 2 && el3 == 3 && q.Empty()) {
		t.Errorf("queue operations disfunctional")
	}
}
