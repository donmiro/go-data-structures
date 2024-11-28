package staks

import "strconv"

type Element struct {
	elementValue int
}

type Stack struct {
	Elements     []*Element
	ElementCount int
}

func (element *Element) String() string {
	return strconv.Itoa(element.elementValue)
}

func (stack *Stack) New() {
	stack.Elements = make([]*Element, 0)
}

// The Push method adds the node to the top of the stack class:
func (stack *Stack) Push(element *Element) {
	stack.Elements = append(stack.Elements[:stack.ElementCount], element)
	stack.ElementCount += 1
}

// The Pop method on the Stack implimentation removes the last element from the element array and returns the element^
func (stack *Stack) Pop() *Element {
	if stack.ElementCount == 0 {
		return nil
	}
	var length int = len(stack.Elements)
	var element *Element = stack.Elements[length-1]

	if length > 1 {
		stack.Elements = stack.Elements[:length-1]
	} else {
		stack.Elements = stack.Elements[0:]
	}
	stack.ElementCount = len(stack.Elements)
	return element
}
