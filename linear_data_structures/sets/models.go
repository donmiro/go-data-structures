package sets

// Set class:
type Set struct {
	integerMap map[int]bool
}

// Create the map of integers and bool:
func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}

// The ContainsElement method of the Set checks element exists or not in the integerMap:
func (set *Set) ContainsElement(element int) bool {
	_, exists := set.integerMap[element]
	return exists
}

// The AddElement method adds the element to a set:
func (set *Set) AddElement(element int) {
	if !set.ContainsElement(element) {
		set.integerMap[element] = true
	}
}

// The DeleteElement method deletes the element from integerMap
func (set *Set) DeleteElement(element int) {
	delete(set.integerMap, element)
}

// The Intersect method return Set that consists of the intersection elements between two sets:
func (set *Set) Intersect(anotherSet *Set) *Set {
	var intersectSet = &Set{}
	intersectSet.New()
	for value := range set.integerMap {
		if anotherSet.ContainsElement(value) {
			intersectSet.AddElement(value)
		}
	}
	return intersectSet
}

// The Union method return  union of two sets:
func (set *Set) Union(anotherSet *Set) *Set {
	var unionSet = &Set{}
	unionSet.New()
	for value := range set.integerMap {
		unionSet.AddElement(value)
	}
	for value := range anotherSet.integerMap {
		unionSet.AddElement(value)
	}
	return unionSet
}
