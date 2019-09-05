package list

import (
	"testing"
)

func TestAddToSingleList(t *testing.T) {
	list := NewSingleList()
	list.Add(12)
	list.Add(1)

	node1 := list.Head
	node2 := list.Head.Next

	if node1.Value != 12 {
		t.Fatalf("list.Head.Value should be 12 but actual is %v", list.Head.Value)
	}
	if node2.Value != 1 {
		t.Fatalf("list.Head.Value should be 1 but actual is %v", list.Head.Value)
	}
}

func TestAddNodeToSingleList(t *testing.T) {
	list := NewSingleList()
	list.AddNode(NewNode(12))
	list.AddNode(NewNode(1))

	node1 := list.Head
	node2 := list.Head.Next

	if node1.Value != 12 {
		t.Fatalf("list.Head.Value should be 12 but actual is %v", list.Head.Value)
	}
	if node2.Value != 1 {
		t.Fatalf("list.Head.Value should be 1 but actual is %v", list.Head.Value)
	}
}
