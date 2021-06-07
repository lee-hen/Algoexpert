package topological_sort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	jobs := jobs(4)
	deps := []Dep{{1, 2}, {1, 3}, {3, 2}, {4, 2}, {4, 3}}
	order := TopologicalSort(jobs, deps)
	if !isValidTopologicalOrder(order, jobs, deps) {
		t.Fail()
	}
}

func TestCase2(t *testing.T) {
	jobs := jobs1(12)
	deps := []Dep{
		{1, 2},
		{1, 3},
		{1, 4},
		{1, 5},
		{1, 6},
		{1, 7},
		{2, 8},
		{3, 8},
		{4, 8},
		{5, 8},
		{6, 8},
		{7, 8},
		{2, 3},
		{2, 4},
		{5, 4},
		{7, 6},
		{6, 2},
		{6, 3},
		{6, 5},
		{5, 9},
		{9, 8},
		{8, 0},
		{4, 0},
		{5, 0},
		{9, 0},
		{2, 0},
		{3, 9},
		{3, 10},
		{10, 11},
		{11, 12},
		{12, 2},
	}
	order := TopologicalSort(jobs, deps)
	require.Equal(t, []int{}, order)
}

func TestCase3(t *testing.T) {
	jobs := jobs(8)
	deps := []Dep{
		{3, 1},
		{8, 1},
		{8, 7},
		{5, 7},
		{5, 2},
		{1, 4},
		{1, 6},
		{1, 2},
		{7, 6},
	}
	order := TopologicalSort(jobs, deps)
	if !isValidTopologicalOrder(order, jobs, deps) {
		t.Fail()
	}
}

func TestCase4(t *testing.T) {
	jobs := jobs(8)
	deps := []Dep{
		{3, 1},
		{8, 1},
		{8, 7},
		{5, 7},
		{5, 2},
		{1, 4},
		{6, 7},
		{1, 2},
		{7, 6},
	}
	order := TopologicalSort(jobs, deps)
	require.Equal(t, []int{}, order)
}

func TestCase5(t *testing.T) {
	jobs := jobs(8)
	deps := []Dep{
		{3, 1},
		{8, 1},
		{8, 7},
		{5, 7},
		{5, 2},
		{1, 4},
		{1, 6},
		{1, 2},
		{7, 6},
		{4, 6},
		{6, 2},
		{2, 3},
	}
	order := TopologicalSort(jobs, deps)
	require.Equal(t, []int{}, order)
}

func jobs1(n int) []int {
	out := []int{}
	for i := 0; i < n; i++ {
		out = append(out, i)
	}
	return out
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
			if _, found := visited[dep.Job]; found && candidate == dep.PreReq {
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
