package ds

type Set []int

func NewSet(numEls int) *Set {
	s := Set(make([]int, numEls+1))
	return &s
}

func (s *Set) QuerySet(x int) int {
	if (*s)[x] == 0 {
		return x
	}
	(*s)[x] = s.QuerySet((*s)[x])
	return (*s)[x]
}

func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (s *Set) Union(x, y int) {
	xS, yS := s.QuerySet(x), s.QuerySet(y)
	xS, yS = MinInt(xS, yS), MaxInt(xS, yS)
	(*s)[xS] = yS
}

func (s *Set) InSameSet(x, y int) bool {
	xS, yS := s.QuerySet(x), s.QuerySet(y)
	return xS == yS
}
