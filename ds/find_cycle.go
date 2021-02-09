package ds

type State int

const (
	S State = iota
	P
	C
)

func isCycle(u int, s []State, g [][]int) bool {
	if s[u] == C {
		return false
	}
	if s[u] == P {
		return true
	}
	s[u] = P
	cycle := false
	for _, v := range g[u] {
		cycle = cycle || isCycle(v, s, g)
		if cycle {
			return true
		}
	}
	s[u] = C
	return false
}

// https://leetcode.com/problems/course-schedule
func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make([][]int, numCourses)
	s := make([]State, numCourses)
	for _, e := range prerequisites {
		v, u := e[0], e[1]
		g[u] = append(g[u], v)
	}
	cycle := false
	for u := 0; u < numCourses; u++ {
		cycle = cycle || isCycle(u, s, g)
		if cycle {
			return false
		}
	}
	return true
}
