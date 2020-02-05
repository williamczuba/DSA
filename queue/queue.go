package queue

import (
	"container/list"

	"github.com/williamczuba/DSA/otherds"
)

type ListQueue struct {
	q *list.List
}

func New() otherds.NodeQueue {
	lq := &ListQueue{list.New()}
	return lq
}

func (l *ListQueue) Enqueue(node *otherds.Node) {
	l.q.PushBack(node)
}

func (l *ListQueue) Dequeue() *otherds.Node {
	front := l.q.Front()
	l.q.Remove(front)
	node, _ := front.Value.(*otherds.Node)
	return node
}

func (l *ListQueue) Len() int {
	return l.q.Len()
}

func (l *ListQueue) IsEmpty() bool {
	return l.q.Len() == 0
}
