package bst

type (
	API interface {
		Insert(value int)
	}

	Node struct {
		Value int
		Left  *Node
		Right *Node
	}

	BST struct {
		Root *Node
	}
)

func (bst *BST) Insert(value int) {
	node := &Node{value, nil, nil}
	if bst.Root == nil {
		bst.Root = node
	} else {
		bst.Root.InsertNode(node)
	}
}

func (n *Node) InsertNode(node *Node) {
	if node.Value < n.Value {
		if n.Left == nil {
			n.Left = node
		} else {
			n.Left.InsertNode(node)
		}
	} else {
		if n.Right == nil {
			n.Right = node
		} else {
			n.Right.InsertNode(node)
		}
	}
}

func Contains(value int, node *Node) bool {
	// check if nil root/ is empty
	if node == nil {
		return false
	}
	// check if this node is the value
	if node.Value == value {
		return true
	}
	// else go left or right.
	if value < node.Value {
		return Contains(value, node.Left)
	} else {
		return Contains(value, node.Right)
	}
}
