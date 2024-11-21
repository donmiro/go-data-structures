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


