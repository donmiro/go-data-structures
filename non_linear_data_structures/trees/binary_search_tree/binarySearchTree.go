package binarysearchtree

import (
	"fmt"
	"sync"
)

// Tree node class:
type TreeNode struct {
	key       int
	value     int
	leftNode  *TreeNode
	rightNode *TreeNode
}

// BinarySearchTree class:
type BinarySearchTree struct {
	rootNode *TreeNode
	lock     sync.RWMutex
}

// InsertElement method:
func (tree *BinarySearchTree) InsertElement(key, value int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	treeNode := &TreeNode{key, value, nil, nil}
	if tree.rootNode == nil {
		tree.rootNode = treeNode
	} else {
		insertTreeNode(tree.rootNode, treeNode)
	}
}

// InOrderTraverseTree method:
func (tree *BinarySearchTree) InOrderTraverseTree(function func(int)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	inOrderTraverseTree(tree.rootNode, function)
}

// PreOrderTraverseTree method:
func (tree *BinarySearchTree) PreOrderTraverseTree(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	preOrderTraverseTree(tree.rootNode, function)
}

// MinNode method:
func (tree *BinarySearchTree) MinNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	treeNode := tree.rootNode
	if treeNode == nil {
		return (*int)(nil)
	}
	for {
		if treeNode.leftNode == nil {
			return &treeNode.value
		}
		treeNode = treeNode.leftNode
	}
}

// MaxNode method:
func (tree *BinarySearchTree) MaxNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	treeNode := tree.rootNode
	if treeNode == nil {
		return (*int)(nil)
	}
	for {
		if treeNode.rightNode == nil {
			return &treeNode.value
		}
		treeNode = treeNode.rightNode
	}
}

// PostOrderTraverseTree method:
func (tree *BinarySearchTree) PostOrderTraverseTree(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	postOrderTraverseTree(tree.rootNode, function)
}

// SearchNode method:
func (tree *BinarySearchTree) SearchNode(key int) bool {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	return searchNode(tree.rootNode, key)
}

// RemoveNode method:
func (tree *BinarySearchTree) RemoveNode(key int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	removeNode(tree.rootNode, key)
}

// String method:
func (tree *BinarySearchTree) String() {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	fmt.Println("--------------------------------------")
	stringify(tree.rootNode, 0)
	fmt.Println("--------------------------------------")
}

// insertTreeNode function:
func insertTreeNode(rootNode, newTreeNode *TreeNode) {
	if newTreeNode.key < rootNode.key {
		if rootNode.leftNode == nil {
			rootNode.leftNode = newTreeNode
		} else {
			insertTreeNode(rootNode.leftNode, newTreeNode)
		}
	} else {
		if rootNode.rightNode == nil {
			rootNode.rightNode = newTreeNode
		} else {
			insertTreeNode(rootNode.rightNode, newTreeNode)
		}
	}
}

// inOrderTraverseTree function:
func inOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		inOrderTraverseTree(treeNode.leftNode, function)
		function(treeNode.value)
		inOrderTraverseTree(treeNode.rightNode, function)
	}
}

// preOrderTraverseTree furnction:
func preOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		function(treeNode.value)
		preOrderTraverseTree(treeNode.leftNode, function)
		preOrderTraverseTree(treeNode.rightNode, function)
	}
}

// preOrderTraverseTree furnction:
func postOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		postOrderTraverseTree(treeNode.leftNode, function)
		postOrderTraverseTree(treeNode.rightNode, function)
		function(treeNode.value)
	}
}

// searckNode function:
func searchNode(treeNode *TreeNode, key int) bool {
	if treeNode == nil {
		return false
	}
	if key < treeNode.key {
		return searchNode(treeNode.leftNode, key)
	}
	if key > treeNode.key {
		return searchNode(treeNode.rightNode, key)
	}
	return true
}

// removeNode function:
func removeNode(treeNode *TreeNode, key int) *TreeNode {
	if treeNode == nil {
		return nil
	}

	if key < treeNode.key {
		treeNode.leftNode = removeNode(treeNode.leftNode, key)
		return treeNode
	}

	if key > treeNode.key {
		treeNode.rightNode = removeNode(treeNode.rightNode, key)
		return treeNode
	}

	// key = node.key
	if treeNode.leftNode == nil && treeNode.rightNode == nil {
		treeNode = nil
		return nil
	}
	if treeNode.leftNode == nil {
		treeNode = treeNode.rightNode
		return treeNode
	}
	if treeNode.rightNode == nil {
		treeNode = treeNode.leftNode
		return treeNode
	}
	leftmostrightNode := treeNode.rightNode
	for {
		// find smallest value on the right side
		if leftmostrightNode != nil && leftmostrightNode.leftNode != nil {
			leftmostrightNode = leftmostrightNode.leftNode
		} else {
			break
		}
	}
	treeNode.key, treeNode.value = leftmostrightNode.key, leftmostrightNode.value
	treeNode.rightNode = removeNode(treeNode.rightNode, treeNode.key)
	return treeNode
}

// Stringify function:
func stringify(treeNode *TreeNode, level int) {
	if treeNode != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += " "
		}
		format += "---[ "
		level++
		stringify(treeNode.leftNode, level)
		fmt.Printf(format+"%d\n", treeNode.key)
		stringify(treeNode.rightNode, level)
	}
}
