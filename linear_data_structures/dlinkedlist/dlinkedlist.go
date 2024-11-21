package dlinkedlist

// Node class:
type Node struct {
	Property     int
	nextNode     *Node
	previousNode *Node
}

// LinkedList class:
type DLinkedList struct {
	HeadNode *Node
}

// The LastNode method returns the node at the end of the list:
func (linkedlist *DLinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node

	for node = linkedlist.HeadNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}

	return lastNode
}

// The NodeBetweenValues method returns the node that has a property lying between values:
func (linkedList *DLinkedList) NodeBetweenValues(firstProperty, secondProperty int) *Node {
	var node, nodeWith *Node
	for node = linkedList.HeadNode; node != nil; node = node.nextNode {
		if node.previousNode != nil && node.nextNode != nil {
			if node.previousNode.Property == firstProperty && node.nextNode.Property == secondProperty {
				nodeWith = node
				break
			}
		}
	}
	return nodeWith
}

// The AddToHead method sets the previousNode property of the current headNode:
func (linkedList *DLinkedList) AddToHead(property int) {
	var node = &Node{}
	node.Property = property
	node.nextNode = nil

	if linkedList.HeadNode != nil {
		node.nextNode = linkedList.HeadNode
		linkedList.HeadNode.previousNode = node
	}
	linkedList.HeadNode = node
}

func (linkedList *DLinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.Property = property
	node.nextNode = nil
	lastNode := linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
		node.previousNode = lastNode
	}
}

func (linkedList *DLinkedList) NodeWithValue(property int) *Node {
	var nodeWith *Node

	for node := linkedList.HeadNode; node != nil; node = node.nextNode {
		if node.Property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

func (linkedList *DLinkedList) AddAfter(nodeProperty int, property int) {
	var node = &Node{}
	var nodeWith *Node
	node.Property = property
	node.nextNode = nil
	nodeWith = linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		node.previousNode = nodeWith
		nodeWith.nextNode = node
	}
}
