package otherds

type NodeQueue interface {
	Enqueue(node *Node)
	Dequeue() *Node
	Len() int
	IsEmpty() bool
}

type Node struct {
	Value     int
	Adjacents []*Node
	Visited   bool
}

type Search struct {
	Q NodeQueue
}

func (s Search) BFS(root *Node, visit func(node *Node)) {
	s.Q.Enqueue(root)
	root.Visited = true
	for !s.Q.IsEmpty() {
		node := s.Q.Dequeue()
		visit(node)

		for _, sib := range node.Adjacents {
			if !sib.Visited {
				sib.Visited = true
				s.Q.Enqueue(sib)
			}
		}
	}
}
