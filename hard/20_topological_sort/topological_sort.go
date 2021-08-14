package topological_sort
// important question

type Dep struct {
	PreReq int
	Job    int
}

type JobNode struct {
	Job      int
	PreReqs  []*JobNode
	Visited  bool
	Visiting bool

	Deps         []*JobNode
	NumOfPreReqs int
}

type JobGraph struct {
	Nodes []*JobNode

	// match job to their node
	Graph map[int]*JobNode
}

func (g *JobGraph) AddNode(job int) {
	g.Graph[job] = &JobNode{Job: job}
	g.Nodes = append(g.Nodes, g.Graph[job])
}

func (g *JobGraph) GetNode(job int) *JobNode {
	if _, found := g.Graph[job]; !found {
		g.AddNode(job)
	}

	return g.Graph[job]
}

// AddPreReq
// depend job point to preReq job
func (g *JobGraph) AddPreReq(job, preReq int) {
	jobNode := g.GetNode(job)
	preReqNode := g.GetNode(preReq)
	jobNode.PreReqs = append(jobNode.PreReqs, preReqNode)
}

// AddDep
// job -> PreReqs
// dep -> job
// preReq job point to depended job
func (g *JobGraph) AddDep(job, dep int) {
	jobNode, depNode := g.GetNode(job), g.GetNode(dep)
	jobNode.Deps = append(jobNode.Deps, depNode)
	depNode.NumOfPreReqs++
}

// TopologicalSort
// O(j + d) time | O(j + d) space
func TopologicalSort(jobs []int, deps []Dep) []int {
	jobGraph := createJobGraph(jobs, deps)
	return getOrderedJobs(jobGraph)
}

// TopologicalSort1
// O(j + d) time | O(j + d) space
func TopologicalSort1(jobs []int, deps []Dep) []int {
	jobGraph := createJobGraph1(jobs, deps)
	return getOrderedJobs1(jobGraph)
}

func createJobGraph(jobs []int, deps []Dep) *JobGraph {
	graph := NewJobGraph(jobs)
	for _, dep := range deps {
		graph.AddPreReq(dep.Job, dep.PreReq)
	}
	return graph
}

func createJobGraph1(jobs []int, deps []Dep) *JobGraph {
	graph := NewJobGraph(jobs)
	for _, dep := range deps {
		graph.AddDep(dep.PreReq, dep.Job)
	}
	return graph
}

func NewJobGraph(jobs []int) *JobGraph {
	g := &JobGraph{
		Graph: map[int]*JobNode{},
	}

	for _, job := range jobs {
		g.AddNode(job)
	}

	return g
}

func getOrderedJobs(graph *JobGraph) []int {
	orderedJobs := make([]int, 0)
	nodes := graph.Nodes

	for len(nodes) != 0 {
		node := nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		containsCycle := depthFirstTraverse(node, &orderedJobs)

		if containsCycle {
			return []int{}
		}
	}

	return orderedJobs
}

func getOrderedJobs1(graph *JobGraph) []int {
	orderedJobs := make([]int, 0)
	nodesWithNoPreReqs := make([]*JobNode, 0)

	for _, node := range graph.Nodes {
		if node.NumOfPreReqs == 0 {
			nodesWithNoPreReqs = append(nodesWithNoPreReqs, node)
		}
	}

	for len(nodesWithNoPreReqs) > 0 {
		node := nodesWithNoPreReqs[len(nodesWithNoPreReqs)-1]
		nodesWithNoPreReqs = nodesWithNoPreReqs[:len(nodesWithNoPreReqs)-1]
		orderedJobs = append(orderedJobs, node.Job)

		removeDeps(node, &nodesWithNoPreReqs)
	}

	for _, node := range graph.Nodes {
		if node.NumOfPreReqs > 0 {
			return []int{}
		}
	}

	return orderedJobs
}

func depthFirstTraverse(node *JobNode, orderedJobs *[]int) bool {
	if node.Visited {
		return false
	} else if node.Visiting {
		return true
	}

	node.Visiting = true
	for _, preReqNode := range node.PreReqs {
		containsCycle := depthFirstTraverse(preReqNode, orderedJobs)
		if containsCycle {
			return true
		}
	}

	node.Visited = true
	node.Visiting = false
	*orderedJobs = append(*orderedJobs, node.Job)

	return false
}

func removeDeps(node *JobNode, nodesWithNoPreReqs *[]*JobNode) {
	for len(node.Deps) > 0 {
		dep := node.Deps[len(node.Deps)-1]
		node.Deps = node.Deps[:len(node.Deps)-1]
		dep.NumOfPreReqs--

		if dep.NumOfPreReqs == 0 {
			*nodesWithNoPreReqs = append(*nodesWithNoPreReqs, dep)
		}
	}
}

// 1 2 3 4
// [1, 2], [1, 3], [3, 2], [4, 2], [4, 3]

// [2:[1] 1:[2]]

// [0:[8 4 5 9 2] 2:[1 6 12] 3:[1 2 6] 4:[1 2 5] 5:[1 6] 6:[1 7] 7:[1] 8:[2 3 4 5 6 6 7 9] 9:[5 3] 10:[3] 11:[10] 12:[11]]

type jobStatus int

const (
	waiting     jobStatus = 1 // unused
	inProgress jobStatus = 2
	completed  jobStatus = 3
	failed     jobStatus = 99
)

// my solution
func topologicalSort(jobs []int, deps []Dep) []int {
	depMap := make(map[int][]int)

	for _, dep := range deps {
		depMap[dep.Job] = append(depMap[dep.Job], dep.PreReq)
	}

	consumedJobs := make([]int, 0)
	visits := make(map[int]jobStatus)

	for _, job := range jobs {
		if visitedJob(job, visits) {
			continue
		}

		if !consumeJobs(job, visits, depMap, &consumedJobs) {
			return []int{}
		}
	}

	return consumedJobs
}

func consumeJobs(job int, visits map[int]jobStatus, depMap map[int][]int, consumedJobs *[]int) bool {
	if consumedJob(job, visits) {
		return true
	}

	if preReqs, found := depMap[job]; found {
		for _, preReq := range preReqs {
			if failedJob(job, visits) {
				return false
			}

			if dependOnEachOther(preReq, depMap, visits) {
				visits[preReq] = failed
				return false
			}

			visits[job] = inProgress
			if !consumeJobs(preReq, visits, depMap, consumedJobs) {
				return false
			}
		}
	}

	*consumedJobs = append(*consumedJobs, job)
	visits[job] = completed

	return true
}

func visitedJob(job int, visits map[int]jobStatus) bool {
	if _, visited := visits[job]; visited {
		return true
	}

	return false
}

func consumedJob(job int, visits map[int]jobStatus) bool {
	if status, visited := visits[job]; visited && status == completed {
		return true
	}

	return false
}

func dependOnEachOther(job int, depMap map[int][]int, visits map[int]jobStatus) bool {
	if _, depend := depMap[job]; depend {
		if status, visited := visits[job]; visited && status == inProgress {
			return true
		}
	}
	return false
}

func failedJob(job int, visits map[int]jobStatus) bool {
	if status, visited := visits[job]; visited && status == failed {
		return true
	}

	return false
}
