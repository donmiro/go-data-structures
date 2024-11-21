package linkedlist

import (
	"fmt"
)

// Node class:
type Node struct {
	Property int
	nextNode *Node
}

// LinkedList class:
type LinkedList struct {
	HeadNode *Node
}

// Methods:

// The AddToHead method adds the node to the start of the linked list:
func (linkedList *LinkedList) AddToHead(property int) {
	var node = Node{}
	node.Property = property
	node.nextNode = linkedList.HeadNode
	linkedList.HeadNode = &node
}

// The IterateList method iterates over LinkedList:
func (linkedList *LinkedList) IterateList() {
	var node *Node
	for node = linkedList.HeadNode; node != nil; node = node.nextNode {
		fmt.Println(node.Property)
	}
}

// The LastNode method returns the node at the end of the list:
func (linkedlist *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node

	for node = linkedlist.HeadNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}

	return lastNode
}

// The AddToEnd method adds the node to the end of the list:
func (linkedList *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.Property = property
	node.nextNode = nil
	lastNode := linkedList.LastNode()
	lastNode.nextNode = node
}

// The NodeWithValue method return the node with the property value:
func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	var nodeWith *Node

	for node := linkedList.HeadNode; node != nil; node = node.nextNode {
		if node.Property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

// The AddAfter method adds the node after the specific node:
func (linkedList *LinkedList) AddAfter(nodeProperty, property int) {
	var node = &Node{}
	var nodeWith *Node

	node.Property = property
	node.nextNode = nil

	nodeWith = linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}
}
