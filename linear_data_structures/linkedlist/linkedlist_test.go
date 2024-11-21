package linkedlist

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestAddToHead(t *testing.T) {
	ll := LinkedList{}
	ll.AddToHead(10)
	ll.AddToHead(20)
	ll.AddToHead(30)

	if ll.HeadNode == nil || ll.HeadNode.Property != 30 {
		t.Errorf("expected head property to be 30, got %v", ll.HeadNode.Property)
	}

	secondNode := ll.HeadNode.nextNode
	if secondNode == nil || secondNode.Property != 20 {
		t.Errorf("expected second node property to be 20, got %v", secondNode.Property)
	}

	thirdNode := secondNode.nextNode
	if thirdNode == nil || thirdNode.Property != 10 {
		t.Errorf("expected third node property to be 10, got %v", thirdNode.Property)
	}

	if thirdNode.nextNode != nil {
		t.Errorf("expected third node next to be nil, got %v", thirdNode.nextNode)
	}
}

func TestIterateList(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.AddToHead(10)
	linkedList.AddToHead(20)
	linkedList.AddToHead(30)

	expectedOutput := "30\n20\n10\n"
	output := captureOutput(linkedList.IterateList)

	if output != expectedOutput {
		t.Errorf("Expected output:\n%sGot:\n%s", expectedOutput, output)
	}
}

func TestIterateListEmpty(t *testing.T) {
	linkedList := LinkedList{}

	expectedOutput := ""
	output := captureOutput(linkedList.IterateList)

	if output != expectedOutput {
		t.Errorf("Expected empty output but got:\n%s", output)
	}
}

func TestLastNode(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.AddToHead(10)
	linkedList.AddToHead(20)
	linkedList.AddToHead(30)

	lastNode := linkedList.LastNode()

	if lastNode == nil {
		t.Fatalf("Expected last node, but got nil")
	}

	if lastNode.Property != 10 {
		t.Errorf("Expected last node property to be 10, but got %d", lastNode.Property)
	}
}

func TestLastNodeEmptyList(t *testing.T) {
	linkedList := LinkedList{}

	lastNode := linkedList.LastNode()

	if lastNode != nil {
		t.Errorf("Expected nil for last node, but got %v", lastNode)
	}
}

func TestAddToEnd(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.AddToHead(10)
	linkedList.AddToEnd(20)
	linkedList.AddToEnd(30)

	firstNode := linkedList.HeadNode
	if firstNode == nil || firstNode.Property != 10 {
		t.Errorf("Expected first node property to be 10, got %v", firstNode.Property)
	}

	secondNode := firstNode.nextNode
	if secondNode == nil || secondNode.Property != 20 {
		t.Errorf("Expected second node property to be 20, got %v", secondNode.Property)
	}

	thirdNode := secondNode.nextNode
	if thirdNode == nil || thirdNode.Property != 30 {
		t.Errorf("Expected third node property to be 30, got %v", thirdNode.Property)
	}

	if thirdNode.nextNode != nil {
		t.Errorf("Expected third node next to be nil, got %v", thirdNode.nextNode)
	}
}

func TestNodeWithValue(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.AddToHead(10)
	linkedList.AddToEnd(20)
	linkedList.AddToEnd(30)

	node := linkedList.NodeWithValue(20)
	if node == nil {
		t.Fatalf("Expected node with value 20, but got nil")
	}

	if node.Property != 20 {
		t.Errorf("Expected node property to be 20, but got %d", node.Property)
	}

	nonExistentNode := linkedList.NodeWithValue(4)
	if nonExistentNode != nil {
		t.Errorf("Expected nil for non-existent node, but got %v", nonExistentNode)
	}
}

func TestAddAfter(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToEnd(3)
	linkedList.AddAfter(1, 2)

	firstNode := linkedList.HeadNode
	if firstNode == nil || firstNode.Property != 1 {
		t.Errorf("Expected first node property to be 1, got %v", firstNode.Property)
	}

	secondNode := firstNode.nextNode
	if secondNode == nil || secondNode.Property != 2 {
		t.Errorf("Expected second node property to be 2, got %v", secondNode.Property)
	}

	thirdNode := secondNode.nextNode
	if thirdNode == nil || thirdNode.Property != 3 {
		t.Errorf("Expected third node property to be 3, got %v", thirdNode.Property)
	}

	if thirdNode.nextNode != nil {
		t.Errorf("Expected third node next to be nil, got %v", thirdNode.nextNode)
	}
}

// Helper functions:
func captureOutput(f func()) string {
	var buf bytes.Buffer
	writer := os.Stdout
	readPipe, writePipe, _ := os.Pipe()
	os.Stdout = writePipe

	f()

	writePipe.Close()
	os.Stdout = writer
	io.Copy(&buf, readPipe)

	return buf.String()
}
