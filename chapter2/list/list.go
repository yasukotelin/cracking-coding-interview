package list

import (
	"encoding/json"
)

type SingleList struct {
	Head   *Node
	Tail   *Node
}

func (sl *SingleList) String() (string, error) {
	bytes, err := json.Marshal(sl)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func NewSingleList() *SingleList {
	return &SingleList{
		Head:   nil,
		Tail:   nil,
	}
}

func (sl *SingleList) Add(value int) {
	node := &Node{Next: nil, Value: value}

	if sl.Head == nil {
		sl.Head = node
		sl.Tail = node
	} else {
		sl.Tail.Next = node
		sl.Tail = node
	}
}

func (sl *SingleList) AddNode(node *Node) {
	if sl.Head == nil {
		sl.Head = node
		sl.Tail = node
	} else {
		sl.Tail.Next = node
		sl.Tail = node
	}
}

type Node struct {
	Next  *Node
	Value int
}

func NewNode(v int) *Node {
	return &Node{
		Next: nil,
		Value: v,
	}
}
