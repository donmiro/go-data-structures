package dlinkedlist

import (
	"testing"
)

func TestAddToHead(t *testing.T) {
	dll := DLinkedList{}
	dll.AddToHead(10)
	dll.AddToHead(20)
	dll.AddToHead(30)

	if dll.HeadNode == nil || dll.HeadNode.Property != 30 {
		t.Errorf("Expected head property to be 30, got %v", dll.HeadNode.Property)
	}

	if dll.HeadNode.nextNode == nil || dll.HeadNode.nextNode.Property != 20 {
		t.Errorf("Expected next property to be 20, got %v", dll.HeadNode.nextNode.Property)
	}

	if dll.HeadNode.nextNode.nextNode == nil || dll.HeadNode.nextNode.nextNode.Property != 10 {
		t.Errorf("Expected next next property to be 10, got %v", dll.HeadNode.nextNode.nextNode.Property)
	}
}

func TestAddToEnd(t *testing.T) {
	dll := DLinkedList{}
	dll.AddToHead(10)
	dll.AddToEnd(20)
	dll.AddToEnd(30)

	lastNode := dll.LastNode()
	if lastNode == nil || lastNode.Property != 30 {
		t.Errorf("Expected last node property to be 30, got %v", lastNode.Property)
	}
}

func TestNodeWithValue(t *testing.T) {
	dll := DLinkedList{}
	dll.AddToHead(10)
	dll.AddToEnd(20)
	dll.AddToEnd(30)

	node := dll.NodeWithValue(20)
	if node == nil || node.Property != 20 {
		t.Errorf("Expected node with property 20, got %v", node.Property)
	}
}

func TestAddAfter(t *testing.T) {
	dll := DLinkedList{}
	dll.AddToHead(10)
	dll.AddToEnd(20)
	dll.AddToEnd(30)
	dll.AddAfter(20, 40)

	node := dll.NodeWithValue(20)
	if node == nil || node.nextNode == nil || node.nextNode.Property != 40 {
		t.Errorf("Expected node with property 4 after node with property 20, got %v", node.nextNode.Property)
	}
}

func TestNodeBetweenValues(t *testing.T) {
	dll := DLinkedList{}
	dll.AddToHead(10)
	dll.AddToEnd(20)
	dll.AddToEnd(30)

	node := dll.NodeBetweenValues(10, 30)
	if node == nil || node.Property != 20 {
		t.Errorf("Expected node with property 20 between values 10 and 30, got %v", node.Property)
	}
}

func TestLastNode(t *testing.T) {
	dll := DLinkedList{}
	dll.AddToHead(10)
	dll.AddToEnd(20)
	dll.AddToEnd(30)

	lastNode := dll.LastNode()
	if lastNode == nil || lastNode.Property != 30 {
		t.Errorf("Expected last node property to be 30, got %v", lastNode.Property)
	}
}
