package staks

import (
	"testing"
)

func TestPush(t *testing.T) {
	stack := Stack{}
	stack.New()

	element1 := &Element{elementValue: 1}
	element2 := &Element{elementValue: 2}

	stack.Push(element1)
	stack.Push(element2)

	if stack.ElementCount != 2 {
		t.Errorf("Expected stack element count to be 2, but got %d", stack.ElementCount)
	}
	if stack.Elements[1] != element2 {
		t.Errorf("Expected element2 to be on top of the stack, but found %v", stack.Elements[1].elementValue)
	}
}

func TestPop(t *testing.T) {
	stack := Stack{}
	stack.New()

	element1 := &Element{elementValue: 1}
	element2 := &Element{elementValue: 2}

	stack.Push(element1)
	stack.Push(element2)

	poppedElement := stack.Pop()

	if poppedElement != element2 {
		t.Errorf("Expected popped element to be element2, but got %v", poppedElement.elementValue)
	}
	if stack.ElementCount != 1 {
		t.Errorf("Expected stack element count to be 1 after pop, but got %d", stack.ElementCount)
	}
	if stack.Elements[0] != element1 {
		t.Errorf("Expected element1 to be on top of the stack after pop, but found %v", stack.Elements[0].elementValue)
	}
}
