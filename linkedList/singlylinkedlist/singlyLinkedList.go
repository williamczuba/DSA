package singlylinkedlist

import "fmt"

import "strconv"

// API is the default integer singly linked list interface (for mocking purposes)
type API interface {
	// Add stores the node in the list
	Add(value int)
	Contains(value int)
	Remove(value int)
	Traverse()
	ReverseTraverse()
	String() string
}

type Node struct {
	Value    int
	nextNode *Node
}

// Int represents an Integer Singly Linked List implementation
type Int struct {
	head *Node
	tail *Node
}

// Add stores the int value in the list in O(1) time
func (i *Int) Add(value int) {
	// create node
	node := &Node{value, nil}
	// check if node should be the head and tail
	if i.head == nil {
		i.head = node
		i.tail = node
		return
	}
	// otherwise,
	// add the new node as the current tail ptr
	i.tail.nextNode = node
	// make the new node the tail
	i.tail = node
}

// String returns a human readable string of the linked list, using -> as a pointer representation.
func (i *Int) String() string {
	// if head is tail, just print it -> nil
	if i.head == i.tail {
		return fmt.Sprintf("%d -> nil", i.head.Value)
	}
	ptr := i.head
	// print head
	str := strconv.Itoa(ptr.Value)
	ptr = ptr.nextNode
	// print middle node values
	for ptr.nextNode != nil {
		str += fmt.Sprintf(" -> %d", ptr.Value)
		ptr = ptr.nextNode
	}

	// print tail value -> nil
	return fmt.Sprintf("%s -> %d -> nil", str, i.tail.Value)
}
