package breadth_first_search

type Node struct {
	Name     string
	Children []*Node
}

// BreadthFirstSearch
//       A
//  B    C    D
//E   F     G   H
//   I  J    K
// Queue B C D
// A
func (n *Node) BreadthFirstSearch(array []string) []string {
	queue := []*Node{n}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		array = append(array, current.Name)
		for _, child := range current.Children {
			queue = append(queue, child)
		}
	}
	return array
}

// my solution
type queue []*Node

func (n *Node) breadthFirstSearch(array []string) []string {
	q := make(queue, 0)
	q.enqueue(n)
	current := q.dequeue()
	for current != nil {
		array = append(array, current.Name)
		for _, curr := range current.Children {
			q.enqueue(curr)
		}
		current = q.dequeue()
	}

	return array
}

func (q *queue) enqueue(n *Node) {
	*q = append(*q, n)
}

func (q *queue) dequeue() *Node {
	qq := *q
	var dq *Node
	if len(*q) > 0 {
		dq = qq[0]
		*q = qq[1:]
	}
	return dq
}
