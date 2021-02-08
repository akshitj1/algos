package ds

import "testing"

func TestSet(t *testing.T) {
	s := NewSet(10)
	s.Union(1, 2)
	s.Union(3, 4)
	if !(s.InSameSet(1, 2) && s.InSameSet(3, 4)) {
		t.Errorf("Set malformed")
	}

	s.Union(1, 3)
	if !s.InSameSet(2, 4) {
		t.Errorf("Set malformed")
	}

}
