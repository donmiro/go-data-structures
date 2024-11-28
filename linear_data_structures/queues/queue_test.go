package queues

import (
	"testing"
)

func TestAdd(t *testing.T) {
	queue := Queue{}

	member1 := &Member{Priority: 1, Content: "Content1", Name: "Member1"}
	member2 := &Member{Priority: 3, Content: "Content2", Name: "Member2"}
	member3 := &Member{Priority: 2, Content: "Content3", Name: "Member3"}

	queue.Add(member1)
	queue.Add(member2)
	queue.Add(member3)

	if queue[0] != member2 {
		t.Errorf("Expected Member2 to be at the front of the queue, but found %v", queue[0].Name)
	}
	if queue[1] != member3 {
		t.Errorf("Expected Member3 to be second in the queue, but found %v", queue[1].Name)
	}
	if queue[2] != member1 {
		t.Errorf("Expected Member1 to be third in the queue, but found %v", queue[2].Name)
	}
}
