package binarysearchtree

import (
	"testing"
)

// TestInsertAndSearch insert and search methods
func TestInsertAndSearch(t *testing.T) {
	tree := &BinarySearchTree{}

	// Insert elements
	tree.InsertElement(50, 50)
	tree.InsertElement(30, 30)
	tree.InsertElement(70, 70)
	tree.InsertElement(20, 20)
	tree.InsertElement(40, 40)
	tree.InsertElement(60, 60)
	tree.InsertElement(80, 80)

	// Test search
	tests := []struct {
		key      int
		expected bool
	}{
		{50, true},
		{30, true},
		{70, true},
		{20, true},
		{90, false}, // This key is absent
	}

	for _, tt := range tests {
		found := tree.SearchNode(tt.key)
		if found != tt.expected {
			t.Errorf("SearchNode(%d) = %v; expected %v", tt.key, found, tt.expected)
		}
	}
}

// TestMinMax test MinNode и MaxNode methods
func TestMinMax(t *testing.T) {
	tree := &BinarySearchTree{}
	tree.InsertElement(50, 50)
	tree.InsertElement(30, 30)
	tree.InsertElement(70, 70)
	tree.InsertElement(20, 20)
	tree.InsertElement(40, 40)
	tree.InsertElement(60, 60)
	tree.InsertElement(80, 80)

	// Test MinNode
	minNode := tree.MinNode()
	if minNode == nil || *minNode != 20 {
		t.Errorf("MinNode() = %v; expected 20", minNode)
	}

	// Test MaxNode
	maxNode := tree.MaxNode()
	if maxNode == nil || *maxNode != 80 {
		t.Errorf("MaxNode() = %v; expected 80", maxNode)
	}
}

// TestRemove test remove method
func TestRemove(t *testing.T) {
	tree := &BinarySearchTree{}
	tree.InsertElement(50, 50)
	tree.InsertElement(30, 30)
	tree.InsertElement(70, 70)
	tree.InsertElement(20, 20)
	tree.InsertElement(40, 40)
	tree.InsertElement(60, 60)
	tree.InsertElement(80, 80)

	// Delete node (leafs)
	tree.RemoveNode(20)
	if tree.SearchNode(20) {
		t.Errorf("RemoveNode(20) failed; 20 is still present")
	}

	// Delete node with one child
	tree.RemoveNode(30)
	if tree.SearchNode(30) {
		t.Errorf("RemoveNode(30) failed; 30 is still present")
	}

	// Delete node with two children
	tree.RemoveNode(50)
	if tree.SearchNode(50) {
		t.Errorf("RemoveNode(50) failed; 50 is still present")
	}
}

// TestInOrder traverse "in-order"
func TestInOrder(t *testing.T) {
	tree := &BinarySearchTree{}
	tree.InsertElement(50, 50)
	tree.InsertElement(30, 30)
	tree.InsertElement(70, 70)
	tree.InsertElement(20, 20)
	tree.InsertElement(40, 40)
	tree.InsertElement(60, 60)
	tree.InsertElement(80, 80)

	expectedOrder := []int{20, 30, 40, 50, 60, 70, 80}
	var result []int

	tree.InOrderTraverseTree(func(value int) {
		result = append(result, value)
	})

	if len(result) != len(expectedOrder) {
		t.Fatalf("InOrderTraverseTree returned incorrect length")
	}

	for i, v := range result {
		if v != expectedOrder[i] {
			t.Errorf("InOrderTraverseTree() failed; got %v, expected %v", result, expectedOrder)
			break
		}
	}
}

func TestString(t *testing.T) {
	tree := &BinarySearchTree{}
	tree.InsertElement(50, 50)
	tree.InsertElement(30, 30)
	tree.InsertElement(70, 70)
	tree.InsertElement(20, 20)
	tree.InsertElement(40, 40)
	tree.InsertElement(60, 60)
	tree.InsertElement(80, 80)

	tree.String()
}
