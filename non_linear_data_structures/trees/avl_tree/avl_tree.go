package avltree

// KeyValue type:
type KeyValue interface {
	LessThan(KeyValue) bool
	EqualTo(KeyValue) bool
}

// TreeNode struct:
type TreeNode struct {
	KeyValue     KeyValue
	BalanceValue int
	LinkedNodes  []*TreeNode
}

// Opposite method:
func opposite(nodeValue int) int {
	return 1 - nodeValue
}

// Single rotation method:
func singleRotation(rootNode *TreeNode, nodeValue int) *TreeNode {
	saveNode := rootNode.LinkedNodes[opposite(nodeValue)]
	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode.LinkedNodes[nodeValue]
	saveNode.LinkedNodes[nodeValue] = rootNode
	return saveNode
}

// Double rotation method:
func doubleRotation(rootNode *TreeNode, nodeValue int) *TreeNode {
	saveNode := rootNode.LinkedNodes[opposite(nodeValue)].LinkedNodes[nodeValue]

	rootNode.LinkedNodes[opposite(nodeValue)].LinkedNodes[nodeValue] = saveNode.LinkedNodes[opposite(nodeValue)]
	saveNode.LinkedNodes[opposite(nodeValue)] = rootNode.LinkedNodes[opposite(nodeValue)]

	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode
	saveNode = rootNode.LinkedNodes[opposite(nodeValue)]

	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode.LinkedNodes[nodeValue]
	saveNode.LinkedNodes[nodeValue] = rootNode

	return saveNode
}
