package topological_sort

import "testing"

func TestCase1(t *testing.T) {
	jobs := jobs(4)
	deps := []Dep{{1, 2}, {1, 3}, {3, 2}, {4, 2}, {4, 3}}
	order := TopologicalSort(jobs, deps)
	if !isValidTopologicalOrder(order, jobs, deps) {
		t.Fail()
	}
}

func jobs(n int) []int {
	out := []int{}
	for i := 1; i <= n; i++ {
		out = append(out, i)
	}
	return out
}

func isValidTopologicalOrder(order []int, jobs []int, deps []Dep) bool {
	visited := map[int]bool{}
	for _, candidate := range order {
		for _, dep := range deps {
			if _, found := visited[dep.Job]; found && candidate == dep.Prereq {
				return false
			}
		}
		visited[candidate] = true
	}
	for _, job := range jobs {
		if _, found := visited[job]; !found {
			return false
		}
	}
	return len(order) == len(jobs)
}
