package datastruct

import "github.com/cheekybits/genny/generic"

type Item generic.Type

type Node struct {
	value Item
	next  *Node
}

func makeNode(i Item, n *Node) *Node {
	return &Node{i, n}
}

func (this *Node) setNext(that *Node) {
	this.next = that
}
