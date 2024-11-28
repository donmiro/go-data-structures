package sets

import (
	"testing"
)

func TestAddElement(t *testing.T) {
	set := Set{}
	set.New()

	// Add element to set:
	set.AddElement(5)

	// Check that element in the set:
	if !set.ContainsElement(5) {
		t.Errorf("Expected element 5 to be in the set, but it was not found")
	}
}

func TestContainsElements(t *testing.T) {
	set := Set{}
	set.New()

	// Add element in the set:
	set.AddElement(5)
	if !set.ContainsElement(5) {
		t.Errorf("Expected element 5 to be in the set, but it was not found")
	}
	if set.ContainsElement(7) {
		t.Errorf("Did not expect element 7 to be in the set, but it was found")
	}
}

func TestDeleteElement(t *testing.T) {
	set := Set{}
	set.New()

	set.AddElement(7)
	set.DeleteElement(7)

	if set.ContainsElement(7) {
		t.Errorf("Expected element 7 to be deleted from the set, but it was found")
	}
}

func TestIntersect(t *testing.T) {
	setA := Set{}
	setA.New()
	setA.AddElement(1)
	setA.AddElement(2)
	setA.AddElement(3)

	setB := Set{}
	setB.New()
	setB.AddElement(2)
	setB.AddElement(3)
	setB.AddElement(4)

	intersectSet := setA.Intersect(&setB)
	if !intersectSet.ContainsElement(2) || !intersectSet.ContainsElement(3) {
		t.Errorf("Expected elements 2 and 3 to be in the intersection set, but they were not found")
	}
	if intersectSet.ContainsElement(1) || intersectSet.ContainsElement(4) {
		t.Errorf("Did not expect elements 1 or 4 to be in the intersection set, but they were found")
	}
}

func TestUnion(t *testing.T) {
	setA := Set{}
	setA.New()
	setA.AddElement(1)
	setA.AddElement(2)

	setB := Set{}
	setB.New()
	setB.AddElement(3)
	setB.AddElement(4)

	unionSet := setA.Union(&setB)
	if !unionSet.ContainsElement(1) || !unionSet.ContainsElement(2) || !unionSet.ContainsElement(3) || !unionSet.ContainsElement(4) {
		t.Errorf("Expected elements 1, 2, 3, and 4 to be in the union set, but they were not found")
	}
}
