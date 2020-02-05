package otherds_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williamczuba/DSA/otherds"
	"github.com/williamczuba/DSA/queue"
)

func TestBFS(t *testing.T) {
	str := ""
	visit := func(node *otherds.Node) {
		str += fmt.Sprintf("%d ", node.Value)
	}
	root := genTree()
	tests := []struct {
		name     string
		treeRoot *otherds.Node
	}{
		{
			"Test run",
			root,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			search := otherds.Search{queue.New()}
			search.BFS(tt.treeRoot, visit)
		})
	}
	assert.Equal(t, "0 1 2 3 ", str)
}

func genTree() *otherds.Node {
	nodes := getNodes(4)
	nodes[0].Adjacents = nodes[1:3]
	nodes[1].Adjacents = []*otherds.Node{nodes[0], nodes[2], nodes[3]}
	nodes[2].Adjacents = []*otherds.Node{nodes[0], nodes[1], nodes[3]}
	nodes[3].Adjacents = []*otherds.Node{nodes[0], nodes[1], nodes[2]}
	return nodes[0]
}

func getNodes(len int) []*otherds.Node {
	nodes := make([]*otherds.Node, len)
	for i := 0; i < len; i++ {
		nodes[i] = &otherds.Node{i, nil, false}
	}
	return nodes
}
